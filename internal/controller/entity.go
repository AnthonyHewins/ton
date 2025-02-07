package controller

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/AnthonyHewins/ton/gen/go/entity/v0"
	"github.com/AnthonyHewins/ton/pkg/ton"
	"github.com/AnthonyHewins/tradovate"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/proto"
)

type EntityPublisher struct {
	prefix string
	js     jetstream.JetStream

	orderKV  jetstream.KeyValue
	orderErr prometheus.Counter

	publishErr prometheus.Counter
	timeout    time.Duration
	logger     *slog.Logger
}

func convertType(t tradovate.EntityType) entity.EntityType {
	switch t {
	case tradovate.EntityTypeAccount:
		return entity.EntityType_ENTITY_TYPE_ACCOUNT
	case tradovate.EntityTypeAccountRiskStatus:
		return entity.EntityType_ENTITY_TYPE_ACCOUNT_RISK_STATUS
	case tradovate.EntityTypeAdminAlert:
		return entity.EntityType_ENTITY_TYPE_ADMIN_ALERT
	case tradovate.EntityTypeAdminAlertSignal:
		return entity.EntityType_ENTITY_TYPE_ADMIN_ALERT_SIGNAL
	case tradovate.EntityTypeCashBalance:
		return entity.EntityType_ENTITY_TYPE_CASH_BALANCE
	case tradovate.EntityTypeCashBalanceLog:
		return entity.EntityType_ENTITY_TYPE_CASH_BALANCE_LOG
	case tradovate.EntityTypeChat:
		return entity.EntityType_ENTITY_TYPE_CHAT
	case tradovate.EntityTypeChatMessage:
		return entity.EntityType_ENTITY_TYPE_CHAT_MESSAGE
	case tradovate.EntityTypeClearingHouse:
		return entity.EntityType_ENTITY_TYPE_CLEARING_HOUSE
	case tradovate.EntityTypeCommand:
		return entity.EntityType_ENTITY_TYPE_COMMAND
	case tradovate.EntityTypeCommandReport:
		return entity.EntityType_ENTITY_TYPE_COMMAND_REPORT
	case tradovate.EntityTypeContactInfo:
		return entity.EntityType_ENTITY_TYPE_CONTACT_INFO
	case tradovate.EntityTypeContract:
		return entity.EntityType_ENTITY_TYPE_CONTRACT
	case tradovate.EntityTypeContractGroup:
		return entity.EntityType_ENTITY_TYPE_CONTRACT_GROUP
	case tradovate.EntityTypeContractMargin:
		return entity.EntityType_ENTITY_TYPE_CONTRACT_MARGIN
	case tradovate.EntityTypeContractMaturity:
		return entity.EntityType_ENTITY_TYPE_CONTRACT_MATURITY
	case tradovate.EntityTypeCurrency:
		return entity.EntityType_ENTITY_TYPE_CURRENCY
	case tradovate.EntityTypeCurrencyRate:
		return entity.EntityType_ENTITY_TYPE_CURRENCY_RATE
	case tradovate.EntityTypeEntitlement:
		return entity.EntityType_ENTITY_TYPE_ENTITLEMENT
	case tradovate.EntityTypeExchange:
		return entity.EntityType_ENTITY_TYPE_EXCHANGE
	case tradovate.EntityTypeExecutionReport:
		return entity.EntityType_ENTITY_TYPE_EXECUTION_REPORT
	case tradovate.EntityTypeFill:
		return entity.EntityType_ENTITY_TYPE_FILL
	case tradovate.EntityTypeFillFee:
		return entity.EntityType_ENTITY_TYPE_FILL_FEE
	case tradovate.EntityTypeFillPair:
		return entity.EntityType_ENTITY_TYPE_FILL_PAIR
	case tradovate.EntityTypeMarginSnapshot:
		return entity.EntityType_ENTITY_TYPE_MARGIN_SNAPSHOT
	case tradovate.EntityTypeMarketDataSubscription:
		return entity.EntityType_ENTITY_TYPE_MARKET_DATA_SUBSCRIPTION
	case tradovate.EntityTypeMarketDataSubscriptionExchangeScope:
		return entity.EntityType_ENTITY_TYPE_MARKET_DATA_SUBSCRIPTION_EXCHANGE_SCOPE
	case tradovate.EntityTypeMarketDataSubscriptionPlan:
		return entity.EntityType_ENTITY_TYPE_MARKET_DATA_SUBSCRIPTION_PLAN
	case tradovate.EntityTypeOrderStrategy:
		return entity.EntityType_ENTITY_TYPE_ORDER_STRATEGY
	case tradovate.EntityTypeOrderStrategyLink:
		return entity.EntityType_ENTITY_TYPE_ORDER_STRATEGY_LINK
	case tradovate.EntityTypeOrderStrategyType:
		return entity.EntityType_ENTITY_TYPE_ORDER_STRATEGY_TYPE
	case tradovate.EntityTypeOrderVersion:
		return entity.EntityType_ENTITY_TYPE_ORDER_VERSION
	case tradovate.EntityTypeOrganization:
		return entity.EntityType_ENTITY_TYPE_ORGANIZATION
	case tradovate.EntityTypePermissionedAccountAutoLiq:
		return entity.EntityType_ENTITY_TYPE_PERMISSIONED_ACCOUNT_AUTO_LIQ
	case tradovate.EntityTypePosition:
		return entity.EntityType_ENTITY_TYPE_POSITION
	case tradovate.EntityTypeProduct:
		return entity.EntityType_ENTITY_TYPE_PRODUCT
	case tradovate.EntityTypeProductMargin:
		return entity.EntityType_ENTITY_TYPE_PRODUCT_MARGIN
	case tradovate.EntityTypeProductSession:
		return entity.EntityType_ENTITY_TYPE_PRODUCT_SESSION
	case tradovate.EntityTypeProperty:
		return entity.EntityType_ENTITY_TYPE_PROPERTY
	case tradovate.EntityTypeSecondMarketDataSubscription:
		return entity.EntityType_ENTITY_TYPE_SECOND_MARKET_DATA_SUBSCRIPTION
	case tradovate.EntityTypeSpreadDefinition:
		return entity.EntityType_ENTITY_TYPE_SPREAD_DEFINITION
	case tradovate.EntityTypeTradingPermission:
		return entity.EntityType_ENTITY_TYPE_TRADING_PERMISSION
	case tradovate.EntityTypeTradovateSubscription:
		return entity.EntityType_ENTITY_TYPE_TRADOVATE_SUBSCRIPTION
	case tradovate.EntityTypeTradovateSubscriptionPlan:
		return entity.EntityType_ENTITY_TYPE_TRADOVATE_SUBSCRIPTION_PLAN
	case tradovate.EntityTypeUser:
		return entity.EntityType_ENTITY_TYPE_USER
	case tradovate.EntityTypeUserAccountAutoLiq:
		return entity.EntityType_ENTITY_TYPE_USER_ACCOUNT_AUTO_LIQ
	case tradovate.EntityTypeUserAccountPositionLimit:
		return entity.EntityType_ENTITY_TYPE_USER_ACCOUNT_POSITION_LIMIT
	case tradovate.EntityTypeUserAccountRiskParameter:
		return entity.EntityType_ENTITY_TYPE_USER_ACCOUNT_RISK_PARAMETER
	case tradovate.EntityTypeUserPlugin:
		return entity.EntityType_ENTITY_TYPE_USER_PLUGIN
	case tradovate.EntityTypeUserProperty:
		return entity.EntityType_ENTITY_TYPE_USER_PROPERTY
	case tradovate.EntityTypeUserSession:
		return entity.EntityType_ENTITY_TYPE_USER_SESSION
	case tradovate.EntityTypeUserSessionStats:
		return entity.EntityType_ENTITY_TYPE_USER_SESSION_STATS
	default:
		return entity.EntityType_ENTITY_TYPE_UNSPECIFIED
	}
}

func eventType(e tradovate.EventType) entity.Event {
	switch e {
	case tradovate.EventTypeCreated:
		return entity.Event_EVENT_CREATED
	case tradovate.EventTypeDeleted:
		return entity.Event_EVENT_DELETED
	case tradovate.EventTypeUpdated:
		return entity.Event_EVENT_UPDATED
	default:
		return entity.Event_EVENT_UNSPECIFIED
	}
}

func (e *EntityPublisher) Publish(data *tradovate.EntityMsg) {
	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	if err := e.pub(ctx, data); err != nil {
		e.publishErr.Inc()
	}
}

func (e *EntityPublisher) pub(ctx context.Context, data *tradovate.EntityMsg) error {
	switch data.Type {
	case tradovate.EntityTypeOrder:
		o, err := data.Order()
		if err != nil {
			e.logger.ErrorContext(ctx, "failed casting order event type as order", "event", data.Event, "type", data.Type, "raw", data.Data)
			return err
		}

		return e.publishOrder(ctx, o)
	}

	buf, err := proto.Marshal(&entity.Entity{
		Type:  convertType(data.Type),
		Event: eventType(data.Event),
		Raw:   data.Data,
	})

	if err != nil {
		e.logger.ErrorContext(ctx, "failed proto marshal of data", "err", err, "event", data.Event, "type", data.Type, "raw", data.Data)
		return err
	}

	_, err = e.js.Publish(ctx,
		fmt.Sprintf(
			"%s.entity.%s.%s",
			e.prefix,
			strings.ToLower(data.Type.String()),
			strings.ToLower(data.Event.String()),
		),
		buf,
	)

	if err != nil {
		return nil
	}

	e.logger.ErrorContext(ctx, "failed publishing message", "err", err)
	return err
}

func (e *EntityPublisher) publishOrder(ctx context.Context, o *tradovate.Order) error {
	buf, err := proto.Marshal(ton.OrderProtoV0(o))
	if err != nil {
		e.logger.ErrorContext(ctx, "failed marshal of order", "order", o)
		return err
	}

	if _, err = e.orderKV.Put(ctx, fmt.Sprintf("order.%d", o.ID), buf); err != nil {
		e.logger.ErrorContext(ctx, "failed putting order into kv", "order", o, "err", err)
		return err
	}

	return nil
}
