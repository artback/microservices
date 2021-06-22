package parser

import (
	"io"
	"micro_services/api/v1/port"
	"reflect"
	"strings"
	"testing"
)

func TestReadPorts(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []port.Port
		wantErr bool
	}{
		{
			name: "decode port without error",
			args: args{
				strings.NewReader(`{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  }
}`),
			},
			want: []port.Port{
				{
					ID:      "AEAJM",
					Name:    "Ajman",
					City:    "Ajman",
					Country: "United Arab Emirates",
					Coordinates: []float32{
						55.513645,
						25.405216,
					},
					Province: "Ajman",
					Timezone: "Asia/Dubai",
					Unlocs: []string{
						"AEAJM",
					},
					Code:    "52000",
					Regions: []string{},
				},
			},
		},
		{
			name: "decode port with ID error",
			args: args{
				strings.NewReader(`{
  : {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  }
}`),
			},
			wantErr: true,
		},
		{
			name: "decode port with Opening error",
			args: args{
				strings.NewReader(`*{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  }
}`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan Result)
			go ReadPorts(tt.args.reader, ch)
			var err error
			var ports []port.Port
			for result := range ch {
				if result.Err != nil {
					err = result.Err
					break
				}
				ports = append(ports, *result.Port)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadPorts() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(ports, tt.want) {
				t.Errorf("ReadPorts() got= %v, want %v", ports, tt.want)
			}
		})
	}
}
