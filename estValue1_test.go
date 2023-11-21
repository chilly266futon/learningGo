package main

import (
	"testing"
)

func TestEstimateValue2(t *testing.T) {
	// args - описывает аргументы тестируемой функции
	type args struct {
		value int
	}
	// описывает структуру тестовых данных и сами тесты
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Small",
			args: args{
				value: 9,
			},
			want:    "small",
			wantErr: false,
		},
		{
			name: "Medium",
			args: args{
				value: 1000,
			},
			want:    "big",
			wantErr: false,
		},
	}
	// вызываем тестируемую функцию для каждого тест-кейса
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EstimateValue(tt.args.value)
			if got != tt.want {
				t.Errorf("EstimateValue() = %v, want %v", got, tt.want)
			}
		})
	}

}
