package streamSvc

import (
	"io"

	pb_svc_stream "github.com/byeol-i/firebase-auth-module/pb/svc/stream"
	pb_unit_common "github.com/byeol-i/firebase-auth-module/pb/unit/common"

	auth "github.com/byeol-i/firebase-auth-module/pkg/authentication/firebase"
	"github.com/byeol-i/firebase-auth-module/pkg/logger"
	"go.uber.org/zap"

	"os"
)

type StreamSrv struct {
	pb_svc_stream.StreamServer
	app *auth.FirebaseApp
}

func NewStreamServiceServer(app *auth.FirebaseApp) *StreamSrv {
	return &StreamSrv{app: app}
}

func (s StreamSrv) VerifyIdToken(stream pb_svc_stream.Stream_VerifyIdTokenServer) error {
	
	var result string
	
	for {
        res, err := stream.Recv()

        if err == io.EOF {
            logger.Info("receive done")
            logger.Info(" ")
            break
        }

        if err != nil {
            logger.Error("err", zap.Any("e", err))
            os.Remove("temp")
            return nil
        }

        logger.Info("receive ", zap.Any("token", res.Token))
	
		result, err = s.app.VerifyIDToken(stream.Context(), res.Token)
		if err != nil {
            logger.Error("err", zap.Any("e", err))
			// return nil
			return stream.Send(&pb_svc_stream.Message{
				Result: &pb_unit_common.ReturnMsg{
					Error: "Can't verify token",
				},
			})
		}
		stream.SendMsg(&pb_svc_stream.Message{
			Result: &pb_unit_common.ReturnMsg{
				Result: result,
			},
		})

    }
	
    return nil
}
