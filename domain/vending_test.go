package domain

import (
	"reflect"
	"testing"
	"vending-mechine/data"
	"vending-mechine/pkg"
)

func TestVendingMachineDomain_BuyCoca(t *testing.T) {

	type fields struct {
		Machines []*pkg.Machine
	}
	type args struct {
		machine string
		coin    int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "coin_amount_is_not_enough",
			fields: fields{
				Machines: data.Mechins,
			},
			args: args{
				machine: "machine_1",
				coin:    9,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "invalid_product_name",
			fields: fields{
				Machines: data.Mechins,
			},
			args: args{
				machine: "unknown_machine_name",
				coin:    20,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VendingMachineDomain{
				Machines: tt.fields.Machines,
			}
			got, err := v.BuyCoca(tt.args.machine, tt.args.coin)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuyCoca() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BuyCoca() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVendingMachineDomain_BuyCoffee(t *testing.T) {
	type fields struct {
		Machines []*pkg.Machine
	}
	type args struct {
		machine string
		coin    int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "coin_amount_is_not_enough",
			fields: fields{
				Machines: data.Mechins,
			},
			args: args{
				machine: "machine_1",
				coin:    19,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "invalid_product_name",
			fields: fields{
				Machines: data.Mechins,
			},
			args: args{
				machine: "unknown_machine_name",
				coin:    20,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VendingMachineDomain{
				Machines: tt.fields.Machines,
			}
			got, err := v.BuyCoffee(tt.args.machine, tt.args.coin)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuyCoffee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BuyCoffee() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVendingMachineDomain_BuyOperation(t *testing.T) {
	type fields struct {
		Machines []*pkg.Machine
	}
	type args struct {
		mch         *pkg.Machine
		coin        int32
		productName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "product_not_available",
			fields: fields{
				Machines: data.Mechins,
			},
			args: args{
				mch:         data.Mechins[0],
				coin:        20,
				productName: data.CoffeeName + "invalid_product_name",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "coin_amount_is_not_enough",
			fields: fields{
				Machines: data.Mechins,
			},
			args: args{
				mch:         data.Mechins[0],
				coin:        3,
				productName: data.CocaName,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VendingMachineDomain{
				Machines: tt.fields.Machines,
			}
			got, err := v.BuyOperation(tt.args.mch, tt.args.coin, tt.args.productName)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuyOperation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BuyOperation() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVendingMachineDomain_Machinist(t *testing.T) {
	type fields struct {
		Machines []*pkg.Machine
	}
	tests := []struct {
		name   string
		fields fields
		want   []*pkg.Machine
	}{
		{
			name: "get_machine_list",
			fields: fields{
				Machines: data.Mechins,
			},
			want: data.Mechins,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VendingMachineDomain{
				Machines: tt.fields.Machines,
			}
			got := v.MachineList()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Machineslist() got = %v, want %v", got, tt.want)
			}
		})
	}
}
