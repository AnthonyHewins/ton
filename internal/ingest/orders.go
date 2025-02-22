package ingest

import (
	"context"
	"errors"
	"log/slog"

	"github.com/AnthonyHewins/tradovate"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func orderErrReasonToCode(x tradovate.OrderErrReason) codes.Code {
	switch x {
	case tradovate.OrderErrReasonAdvancedTrailingStopUnsupported,
		tradovate.OrderErrReasonUnsupported:
		return codes.Unimplemented
	case tradovate.OrderErrReasonExecutionProviderNotConfigured,
		tradovate.OrderErrReasonOtherExecutionRelated:
		return codes.FailedPrecondition
	case tradovate.OrderErrReasonAnotherCommandPending:
		return codes.Aborted
	case tradovate.OrderErrReasonAccountUnspecified,
		tradovate.OrderErrReasonBackMonthProhibited,
		tradovate.OrderErrReasonInvalidContract,
		tradovate.OrderErrReasonInvalidPrice,
		tradovate.OrderErrReasonNoQuote,
		tradovate.OrderErrReasonParentRejected,
		tradovate.OrderErrReasonLiquidationOnly,
		tradovate.OrderErrReasonLiquidationOnlyBeforeExpiration,
		tradovate.OrderErrReasonMaxOrderQtyIsNotSpecified,
		tradovate.OrderErrReasonMaxPosLimitMisconfigured,
		tradovate.OrderErrReasonTrailingStopNonOrderQtyModify:
		return codes.InvalidArgument
	case tradovate.OrderErrReasonExecutionProviderUnavailable,
		tradovate.OrderErrReasonTradingLocked,
		tradovate.OrderErrReasonAccountClosed:
		return codes.Unavailable
	case tradovate.OrderErrReasonMultipleAccountPlanRequired,
		tradovate.OrderErrReasonUnauthorized:
		return codes.PermissionDenied
	case tradovate.OrderErrReasonRiskCheckTimeout,
		tradovate.OrderErrReasonTooLate:
		return codes.DeadlineExceeded
	case tradovate.OrderErrReasonSessionClosed,
		tradovate.OrderErrReasonMaxOrderQtyLimitReached,
		tradovate.OrderErrReasonMaxTotalPosLimitReached,
		tradovate.OrderErrReasonNotEnoughLiquidity,
		tradovate.OrderErrReasonMaxPosLimitReached:
		return codes.ResourceExhausted
	default:
		return codes.Internal
	}
}

func orderErr(err error) error {
	if errors.Is(err, &tradovate.OrderErr{}) {
		x := err.(*tradovate.OrderErr)
		return status.Error(orderErrReasonToCode(x.Reason), x.Text)
	}

	return status.Error(codes.Internal, err.Error())
}

type Orders struct {
	logger                            *slog.Logger
	placedOrder, placedOCO, placedOSO prometheus.Counter
	orderErr, ocoErr, osoErr          prometheus.Counter
	ws                                *tradovate.WS
}

func (o *Orders) PlaceOrder(ctx context.Context, r *tradovate.OrderReq) (uint, error) {
	id, err := o.ws.PlaceOrder(ctx, r)
	if err != nil {
		o.logger.ErrorContext(ctx, "failed create order", "err", err, "req", r)
		o.orderErr.Inc()
		return 0, orderErr(err)
	}

	o.logger.DebugContext(ctx, "created order", "id", id)
	o.placedOrder.Inc()
	return id, nil
}

func (o *Orders) PlaceOCO(ctx context.Context, r *tradovate.OcoReq) (*tradovate.OcoResp, error) {
	resp, err := o.ws.OCO(ctx, r)
	if err != nil {
		o.logger.ErrorContext(ctx, "failed placing oco", "req", r)
		o.ocoErr.Inc()
		return nil, orderErr(err)
	}

	o.logger.DebugContext(ctx, "placed OCO", "resp", resp)
	o.placedOCO.Inc()
	return resp, nil
}

func (o *Orders) PlaceOSO(ctx context.Context, r *tradovate.OsoReq) (*tradovate.OsoResp, error) {
	resp, err := o.ws.OSO(ctx, r)
	if err != nil {
		o.logger.ErrorContext(ctx, "failed placing oco", "req", r)
		o.osoErr.Inc()
		return nil, orderErr(err)
	}

	o.logger.DebugContext(ctx, "placed OCO", "resp", resp)
	o.placedOSO.Inc()
	return resp, nil
}
