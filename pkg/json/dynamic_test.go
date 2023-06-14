package json

import (
	"reflect"
	"testing"
)

func TestJsonGenerator(t *testing.T) {
	type args struct {
		jsonData []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonGenerator(tt.args.jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonGenerator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonGenerator() got = %v, want %v", got, tt.want)
			}
		})
	}
}
