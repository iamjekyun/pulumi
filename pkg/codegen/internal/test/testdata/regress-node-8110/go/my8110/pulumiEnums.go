// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package my8110

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MyEnum string

const (
	MyEnumOne = MyEnum("one")
	MyEnumTwo = MyEnum("two")
)

func (MyEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*MyEnum)(nil)).Elem()
}

func (e MyEnum) ToMyEnumOutput() MyEnumOutput {
	return pulumi.ToOutput(e).(MyEnumOutput)
}

func (e MyEnum) ToMyEnumOutputWithContext(ctx context.Context) MyEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MyEnumOutput)
}

func (e MyEnum) ToMyEnumPtrOutput() MyEnumPtrOutput {
	return e.ToMyEnumPtrOutputWithContext(context.Background())
}

func (e MyEnum) ToMyEnumPtrOutputWithContext(ctx context.Context) MyEnumPtrOutput {
	return MyEnum(e).ToMyEnumOutputWithContext(ctx).ToMyEnumPtrOutputWithContext(ctx)
}

func (e MyEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MyEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MyEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MyEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MyEnumOutput struct{ *pulumi.OutputState }

func (MyEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MyEnum)(nil)).Elem()
}

func (o MyEnumOutput) ToMyEnumOutput() MyEnumOutput {
	return o
}

func (o MyEnumOutput) ToMyEnumOutputWithContext(ctx context.Context) MyEnumOutput {
	return o
}

func (o MyEnumOutput) ToMyEnumPtrOutput() MyEnumPtrOutput {
	return o.ToMyEnumPtrOutputWithContext(context.Background())
}

func (o MyEnumOutput) ToMyEnumPtrOutputWithContext(ctx context.Context) MyEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MyEnum) *MyEnum {
		return &v
	}).(MyEnumPtrOutput)
}

func (o MyEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MyEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MyEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MyEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MyEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MyEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MyEnumPtrOutput struct{ *pulumi.OutputState }

func (MyEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MyEnum)(nil)).Elem()
}

func (o MyEnumPtrOutput) ToMyEnumPtrOutput() MyEnumPtrOutput {
	return o
}

func (o MyEnumPtrOutput) ToMyEnumPtrOutputWithContext(ctx context.Context) MyEnumPtrOutput {
	return o
}

func (o MyEnumPtrOutput) Elem() MyEnumOutput {
	return o.ApplyT(func(v *MyEnum) MyEnum {
		if v != nil {
			return *v
		}
		var ret MyEnum
		return ret
	}).(MyEnumOutput)
}

func (o MyEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MyEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MyEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// MyEnumInput is an input type that accepts MyEnumArgs and MyEnumOutput values.
// You can construct a concrete instance of `MyEnumInput` via:
//
//          MyEnumArgs{...}
type MyEnumInput interface {
	pulumi.Input

	ToMyEnumOutput() MyEnumOutput
	ToMyEnumOutputWithContext(context.Context) MyEnumOutput
}

var myEnumPtrType = reflect.TypeOf((**MyEnum)(nil)).Elem()

type MyEnumPtrInput interface {
	pulumi.Input

	ToMyEnumPtrOutput() MyEnumPtrOutput
	ToMyEnumPtrOutputWithContext(context.Context) MyEnumPtrOutput
}

type myEnumPtr string

func MyEnumPtr(v string) MyEnumPtrInput {
	return (*myEnumPtr)(&v)
}

func (*myEnumPtr) ElementType() reflect.Type {
	return myEnumPtrType
}

func (in *myEnumPtr) ToMyEnumPtrOutput() MyEnumPtrOutput {
	return pulumi.ToOutput(in).(MyEnumPtrOutput)
}

func (in *myEnumPtr) ToMyEnumPtrOutputWithContext(ctx context.Context) MyEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MyEnumPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MyEnumOutput{})
	pulumi.RegisterOutputType(MyEnumPtrOutput{})
}
