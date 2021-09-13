package secretreader

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

type SSMReader struct {
	client ssmiface.SSMAPI
}

func NewSSMReader(region string) (*SSMReader, error) {
	sess, err := Sessions()
	if err != nil {
		return nil, err
	}

	return &SSMReader{
		client: ssm.New(sess, aws.NewConfig().WithRegion(region)),
	}, nil
}

func (s *SSMReader) Get(key string) (string, error) {
	ssmVal, err := s.client.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(key),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return "", err
	}

	return *ssmVal.Parameter.Value, nil
}

func Sessions() (*session.Session, error) {
	sess, err := session.NewSession()
	svc := session.Must(sess, err)
	return svc, err
}
