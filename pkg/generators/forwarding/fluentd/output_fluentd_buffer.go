package fluentd

import (
	"fmt"

	loggingv1 "github.com/openshift/cluster-logging-operator/pkg/apis/logging/v1"
)

const (
	// Buffer size defaults
	defaultOverflowAction = "block"

	// Flush buffer defaults
	defaultFlushThreadCount = "2"
	defaultFlushMode        = "interval"
	defaultFlushInterval    = "1s"

	// Retry buffer to output defaults
	defaultRetryWait        = "1s"
	defaultRetryType        = "exponential_backoff"
	defaultRetryMaxInterval = "60s"
	defaultRetryTimeout     = "60m"

	// Output fluentdForward default
	fluentdForwardOverflowAction = "block"
	fluentdForwardFlushInterval  = "5s"
)

func (olc *outputLabelConf) ChunkLimitSize() string {
	if hasBufferConfig(olc.forwarder) && olc.forwarder.Fluentd.Buffer.ChunkLimitSize != "" {
		return string(olc.forwarder.Fluentd.Buffer.ChunkLimitSize)
	} else {
		size := ""
		switch olc.Target.Type {
		case loggingv1.OutputTypeFluentdForward:
			size = "1m"
		default:
			size = "8m"
		}
		return fmt.Sprintf("\"#{ENV['BUFFER_SIZE_LIMIT'] || '%s'}\"", size)
	}
}

func (olc *outputLabelConf) TotalLimitSize() string {
	if hasBufferConfig(olc.forwarder) && olc.forwarder.Fluentd.Buffer.TotalLimitSize != "" {
		return string(olc.forwarder.Fluentd.Buffer.TotalLimitSize)
	} else {
		return "\"#{ENV['TOTAL_LIMIT_SIZE'] ||  8589934592 }\" #8G"
	}
}

func (olc *outputLabelConf) OverflowAction() string {
	if hasBufferConfig(olc.forwarder) {
		oa := string(olc.forwarder.Fluentd.Buffer.OverflowAction)

		if oa != "" {
			return oa
		}
	}

	switch olc.Target.Type {
	case loggingv1.OutputTypeFluentdForward:
		return fluentdForwardOverflowAction
	default:
		return defaultOverflowAction
	}
}

func (olc *outputLabelConf) FlushThreadCount() string {
	if hasBufferConfig(olc.forwarder) {
		ftc := olc.forwarder.Fluentd.Buffer.FlushThreadCount

		if ftc > 0 {
			return fmt.Sprintf("%d", ftc)
		}
	}

	return defaultFlushThreadCount
}

func (olc *outputLabelConf) FlushMode() string {
	if hasBufferConfig(olc.forwarder) {
		fm := string(olc.forwarder.Fluentd.Buffer.FlushMode)

		if fm != "" {
			return fm
		}
	}

	return defaultFlushMode
}

func (olc *outputLabelConf) FlushInterval() string {
	if hasBufferConfig(olc.forwarder) {
		fi := string(olc.forwarder.Fluentd.Buffer.FlushInterval)

		if fi != "" {
			return fi
		}
	}

	switch olc.Target.Type {
	case loggingv1.OutputTypeFluentdForward:
		return fluentdForwardFlushInterval
	default:
		return defaultFlushInterval
	}
}

func (olc *outputLabelConf) RetryWait() string {
	if hasBufferConfig(olc.forwarder) {
		rw := string(olc.forwarder.Fluentd.Buffer.RetryWait)

		if rw != "" {
			return rw
		}
	}

	return defaultRetryWait
}

func (olc *outputLabelConf) RetryType() string {
	if hasBufferConfig(olc.forwarder) {
		rt := string(olc.forwarder.Fluentd.Buffer.RetryType)

		if rt != "" {
			return rt
		}
	}

	return defaultRetryType
}

func (olc *outputLabelConf) RetryMaxInterval() string {
	if hasBufferConfig(olc.forwarder) {
		rmi := string(olc.forwarder.Fluentd.Buffer.RetryMaxInterval)

		if rmi != "" {
			return rmi
		}
	}

	return defaultRetryMaxInterval
}
func (olc *outputLabelConf) RetryTimeout() string {
	value := defaultRetryTimeout
	if hasBufferConfig(olc.forwarder) && string(olc.forwarder.Fluentd.Buffer.RetryTimeout) != "" {
		value = string(olc.forwarder.Fluentd.Buffer.RetryTimeout)
	}

	return fmt.Sprintf("retry_timeout %s", value)
}

func hasBufferConfig(config *loggingv1.ForwarderSpec) bool {
	return config != nil && config.Fluentd != nil && config.Fluentd.Buffer != nil
}
