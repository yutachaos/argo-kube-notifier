package integration

import "github.com/argoproj-labs/argo-kube-notifier/util"

var log = util.GetLogger()

type NotifierInterface interface {
	SendSuccessNotification(msg ...string) error
	SendWarningNotification(msg ...string) error
	SendFailledNotification(msg ...string) error
	SendInfoNotification(msg ...string) error
}
