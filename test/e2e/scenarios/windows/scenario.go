package windows

import (
	"github.com/microsoft/retina/test/e2e/framework/types"
)

func ValidateWindowsBasicMetric() *types.Scenario {
	name := "Windows Metrics"
	steps := []*types.StepWrapper{
		{
			Step: &ValidateHNSMetric{
				KubeConfigFilePath:       "./test.pem",
				RetinaDaemonSetNamespace: "kube-system",
				RetinaDaemonSetName:      "retina-agent-win",
			},
		},
	}
	return types.NewScenario(name, steps...)
}