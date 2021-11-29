package s2

import (
	"context"
	"fmt"
	"io"

	"github.com/dr-dobermann/srvbus/ms"
	"github.com/google/uuid"
)

// new.....Service has common signature
//   func(ctx context.Context,
//        params ...serviceParams) (ServiceRunner, error)
//
// the context used here is different from the one used on service
// execution. The context is only used in case there is long
// preparation procedure needed to create a service.
// Usually it would be ignored by "_ context.Context" declaration.

// =============================================================================
//    Output Service

// newOutputService returns an output ServiceFunc which
// puts all values form vl into io.Writer w.
func newOutputService(
	_ context.Context,
	w io.Writer,
	vl ...interface{}) (ServiceRunner, error) {
	if w == nil {
		return nil, fmt.Errorf("writer isn't presented for OutputSvc")
	}

	outputSvc := func(_ context.Context) error {
		fmt.Fprint(w, vl...)

		return nil
	}

	return ServiceFunc(outputSvc), nil
}

// =============================================================================
//    Put Messages Service

func newPutMessagesService(
	_ context.Context,
	mSrv *ms.MessageServer,
	queue string,
	sender uuid.UUID,
	msgs ...*ms.Message) (ServiceRunner, error) {

	if mSrv == nil || queue == "" || len(msgs) == 0 {
		return nil, fmt.Errorf("invalid parameter for PutMessage Service : "+
			"mSrv(%p), queue name(%s), msg num(%d)",
			mSrv, queue, len(msgs))
	}

	putMessages := func(_ context.Context) error {

		return mSrv.PutMessages(sender, queue, msgs...)
	}

	return ServiceFunc(putMessages), nil
}
