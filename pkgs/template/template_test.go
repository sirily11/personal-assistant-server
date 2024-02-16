package template

import (
	"sme-demo/internal/config"
	"testing"
)

func TestTemplateClient_Render(t1 *testing.T) {
	type fields struct {
		config config.Config
	}
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Render simple template",
			fields: fields{
				config: config.Config{
					ChatRoom: config.ChatRoom{
						ChatRoomTitle: "Test",
					},
				},
			},
			args: args{
				content: "{{.Config.ChatRoom.ChatRoomTitle}} Data",
			},
			want: "Test Data",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TemplateClient{
				config: tt.fields.config,
			}
			got, err := t.Render(tt.args.content)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("Render() got = %v, want %v", got, tt.want)
			}
		})
	}
}
