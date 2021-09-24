// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExampleEnum string

const (
	ExampleEnumOne = ExampleEnum("one")
	ExampleEnumTwo = ExampleEnum("two")
)

func (ExampleEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*ExampleEnum)(nil)).Elem()
}

func (e ExampleEnum) ToExampleEnumOutput() ExampleEnumOutput {
	return pulumi.ToOutput(e).(ExampleEnumOutput)
}

func (e ExampleEnum) ToExampleEnumOutputWithContext(ctx context.Context) ExampleEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExampleEnumOutput)
}

func (e ExampleEnum) ToExampleEnumPtrOutput() ExampleEnumPtrOutput {
	return e.ToExampleEnumPtrOutputWithContext(context.Background())
}

func (e ExampleEnum) ToExampleEnumPtrOutputWithContext(ctx context.Context) ExampleEnumPtrOutput {
	return ExampleEnum(e).ToExampleEnumOutputWithContext(ctx).ToExampleEnumPtrOutputWithContext(ctx)
}

func (e ExampleEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExampleEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExampleEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExampleEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExampleEnumOutput struct{ *pulumi.OutputState }

func (ExampleEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExampleEnum)(nil)).Elem()
}

func (o ExampleEnumOutput) ToExampleEnumOutput() ExampleEnumOutput {
	return o
}

func (o ExampleEnumOutput) ToExampleEnumOutputWithContext(ctx context.Context) ExampleEnumOutput {
	return o
}

func (o ExampleEnumOutput) ToExampleEnumPtrOutput() ExampleEnumPtrOutput {
	return o.ToExampleEnumPtrOutputWithContext(context.Background())
}

func (o ExampleEnumOutput) ToExampleEnumPtrOutputWithContext(ctx context.Context) ExampleEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExampleEnum) *ExampleEnum {
		return &v
	}).(ExampleEnumPtrOutput)
}

func (o ExampleEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExampleEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExampleEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExampleEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExampleEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExampleEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExampleEnumPtrOutput struct{ *pulumi.OutputState }

func (ExampleEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExampleEnum)(nil)).Elem()
}

func (o ExampleEnumPtrOutput) ToExampleEnumPtrOutput() ExampleEnumPtrOutput {
	return o
}

func (o ExampleEnumPtrOutput) ToExampleEnumPtrOutputWithContext(ctx context.Context) ExampleEnumPtrOutput {
	return o
}

func (o ExampleEnumPtrOutput) Elem() ExampleEnumOutput {
	return o.ApplyT(func(v *ExampleEnum) ExampleEnum {
		if v != nil {
			return *v
		}
		var ret ExampleEnum
		return ret
	}).(ExampleEnumOutput)
}

func (o ExampleEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExampleEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExampleEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ExampleEnumInput is an input type that accepts ExampleEnumArgs and ExampleEnumOutput values.
// You can construct a concrete instance of `ExampleEnumInput` via:
//
//          ExampleEnumArgs{...}
type ExampleEnumInput interface {
	pulumi.Input

	ToExampleEnumOutput() ExampleEnumOutput
	ToExampleEnumOutputWithContext(context.Context) ExampleEnumOutput
}

var exampleEnumPtrType = reflect.TypeOf((**ExampleEnum)(nil)).Elem()

type ExampleEnumPtrInput interface {
	pulumi.Input

	ToExampleEnumPtrOutput() ExampleEnumPtrOutput
	ToExampleEnumPtrOutputWithContext(context.Context) ExampleEnumPtrOutput
}

type exampleEnumPtr string

func ExampleEnumPtr(v string) ExampleEnumPtrInput {
	return (*exampleEnumPtr)(&v)
}

func (*exampleEnumPtr) ElementType() reflect.Type {
	return exampleEnumPtrType
}

func (in *exampleEnumPtr) ToExampleEnumPtrOutput() ExampleEnumPtrOutput {
	return pulumi.ToOutput(in).(ExampleEnumPtrOutput)
}

func (in *exampleEnumPtr) ToExampleEnumPtrOutputWithContext(ctx context.Context) ExampleEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExampleEnumPtrOutput)
}

type ExampleEnumInputEnum string

const (
	ExampleEnumInputEnumOne = ExampleEnumInputEnum("one")
	ExampleEnumInputEnumTwo = ExampleEnumInputEnum("two")
)

func (ExampleEnumInputEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*ExampleEnumInputEnum)(nil)).Elem()
}

func (e ExampleEnumInputEnum) ToExampleEnumInputEnumOutput() ExampleEnumInputEnumOutput {
	return pulumi.ToOutput(e).(ExampleEnumInputEnumOutput)
}

func (e ExampleEnumInputEnum) ToExampleEnumInputEnumOutputWithContext(ctx context.Context) ExampleEnumInputEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExampleEnumInputEnumOutput)
}

func (e ExampleEnumInputEnum) ToExampleEnumInputEnumPtrOutput() ExampleEnumInputEnumPtrOutput {
	return e.ToExampleEnumInputEnumPtrOutputWithContext(context.Background())
}

func (e ExampleEnumInputEnum) ToExampleEnumInputEnumPtrOutputWithContext(ctx context.Context) ExampleEnumInputEnumPtrOutput {
	return ExampleEnumInputEnum(e).ToExampleEnumInputEnumOutputWithContext(ctx).ToExampleEnumInputEnumPtrOutputWithContext(ctx)
}

func (e ExampleEnumInputEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExampleEnumInputEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExampleEnumInputEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExampleEnumInputEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExampleEnumInputEnumOutput struct{ *pulumi.OutputState }

func (ExampleEnumInputEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExampleEnumInputEnum)(nil)).Elem()
}

func (o ExampleEnumInputEnumOutput) ToExampleEnumInputEnumOutput() ExampleEnumInputEnumOutput {
	return o
}

func (o ExampleEnumInputEnumOutput) ToExampleEnumInputEnumOutputWithContext(ctx context.Context) ExampleEnumInputEnumOutput {
	return o
}

func (o ExampleEnumInputEnumOutput) ToExampleEnumInputEnumPtrOutput() ExampleEnumInputEnumPtrOutput {
	return o.ToExampleEnumInputEnumPtrOutputWithContext(context.Background())
}

func (o ExampleEnumInputEnumOutput) ToExampleEnumInputEnumPtrOutputWithContext(ctx context.Context) ExampleEnumInputEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExampleEnumInputEnum) *ExampleEnumInputEnum {
		return &v
	}).(ExampleEnumInputEnumPtrOutput)
}

func (o ExampleEnumInputEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExampleEnumInputEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExampleEnumInputEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExampleEnumInputEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExampleEnumInputEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExampleEnumInputEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExampleEnumInputEnumPtrOutput struct{ *pulumi.OutputState }

func (ExampleEnumInputEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExampleEnumInputEnum)(nil)).Elem()
}

func (o ExampleEnumInputEnumPtrOutput) ToExampleEnumInputEnumPtrOutput() ExampleEnumInputEnumPtrOutput {
	return o
}

func (o ExampleEnumInputEnumPtrOutput) ToExampleEnumInputEnumPtrOutputWithContext(ctx context.Context) ExampleEnumInputEnumPtrOutput {
	return o
}

func (o ExampleEnumInputEnumPtrOutput) Elem() ExampleEnumInputEnumOutput {
	return o.ApplyT(func(v *ExampleEnumInputEnum) ExampleEnumInputEnum {
		if v != nil {
			return *v
		}
		var ret ExampleEnumInputEnum
		return ret
	}).(ExampleEnumInputEnumOutput)
}

func (o ExampleEnumInputEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExampleEnumInputEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExampleEnumInputEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ExampleEnumInputEnumInput is an input type that accepts ExampleEnumInputEnumArgs and ExampleEnumInputEnumOutput values.
// You can construct a concrete instance of `ExampleEnumInputEnumInput` via:
//
//          ExampleEnumInputEnumArgs{...}
type ExampleEnumInputEnumInput interface {
	pulumi.Input

	ToExampleEnumInputEnumOutput() ExampleEnumInputEnumOutput
	ToExampleEnumInputEnumOutputWithContext(context.Context) ExampleEnumInputEnumOutput
}

var exampleEnumInputEnumPtrType = reflect.TypeOf((**ExampleEnumInputEnum)(nil)).Elem()

type ExampleEnumInputEnumPtrInput interface {
	pulumi.Input

	ToExampleEnumInputEnumPtrOutput() ExampleEnumInputEnumPtrOutput
	ToExampleEnumInputEnumPtrOutputWithContext(context.Context) ExampleEnumInputEnumPtrOutput
}

type exampleEnumInputEnumPtr string

func ExampleEnumInputEnumPtr(v string) ExampleEnumInputEnumPtrInput {
	return (*exampleEnumInputEnumPtr)(&v)
}

func (*exampleEnumInputEnumPtr) ElementType() reflect.Type {
	return exampleEnumInputEnumPtrType
}

func (in *exampleEnumInputEnumPtr) ToExampleEnumInputEnumPtrOutput() ExampleEnumInputEnumPtrOutput {
	return pulumi.ToOutput(in).(ExampleEnumInputEnumPtrOutput)
}

func (in *exampleEnumInputEnumPtr) ToExampleEnumInputEnumPtrOutputWithContext(ctx context.Context) ExampleEnumInputEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExampleEnumInputEnumPtrOutput)
}

type ResourceTypeEnum string

const (
	ResourceTypeEnumHaha     = ResourceTypeEnum("haha")
	ResourceTypeEnumBusiness = ResourceTypeEnum("business")
)

func (ResourceTypeEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceTypeEnum)(nil)).Elem()
}

func (e ResourceTypeEnum) ToResourceTypeEnumOutput() ResourceTypeEnumOutput {
	return pulumi.ToOutput(e).(ResourceTypeEnumOutput)
}

func (e ResourceTypeEnum) ToResourceTypeEnumOutputWithContext(ctx context.Context) ResourceTypeEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceTypeEnumOutput)
}

func (e ResourceTypeEnum) ToResourceTypeEnumPtrOutput() ResourceTypeEnumPtrOutput {
	return e.ToResourceTypeEnumPtrOutputWithContext(context.Background())
}

func (e ResourceTypeEnum) ToResourceTypeEnumPtrOutputWithContext(ctx context.Context) ResourceTypeEnumPtrOutput {
	return ResourceTypeEnum(e).ToResourceTypeEnumOutputWithContext(ctx).ToResourceTypeEnumPtrOutputWithContext(ctx)
}

func (e ResourceTypeEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceTypeEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceTypeEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceTypeEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceTypeEnumOutput struct{ *pulumi.OutputState }

func (ResourceTypeEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceTypeEnum)(nil)).Elem()
}

func (o ResourceTypeEnumOutput) ToResourceTypeEnumOutput() ResourceTypeEnumOutput {
	return o
}

func (o ResourceTypeEnumOutput) ToResourceTypeEnumOutputWithContext(ctx context.Context) ResourceTypeEnumOutput {
	return o
}

func (o ResourceTypeEnumOutput) ToResourceTypeEnumPtrOutput() ResourceTypeEnumPtrOutput {
	return o.ToResourceTypeEnumPtrOutputWithContext(context.Background())
}

func (o ResourceTypeEnumOutput) ToResourceTypeEnumPtrOutputWithContext(ctx context.Context) ResourceTypeEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceTypeEnum) *ResourceTypeEnum {
		return &v
	}).(ResourceTypeEnumPtrOutput)
}

func (o ResourceTypeEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceTypeEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceTypeEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceTypeEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceTypeEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceTypeEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceTypeEnumPtrOutput struct{ *pulumi.OutputState }

func (ResourceTypeEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceTypeEnum)(nil)).Elem()
}

func (o ResourceTypeEnumPtrOutput) ToResourceTypeEnumPtrOutput() ResourceTypeEnumPtrOutput {
	return o
}

func (o ResourceTypeEnumPtrOutput) ToResourceTypeEnumPtrOutputWithContext(ctx context.Context) ResourceTypeEnumPtrOutput {
	return o
}

func (o ResourceTypeEnumPtrOutput) Elem() ResourceTypeEnumOutput {
	return o.ApplyT(func(v *ResourceTypeEnum) ResourceTypeEnum {
		if v != nil {
			return *v
		}
		var ret ResourceTypeEnum
		return ret
	}).(ResourceTypeEnumOutput)
}

func (o ResourceTypeEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceTypeEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceTypeEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ResourceTypeEnumInput is an input type that accepts ResourceTypeEnumArgs and ResourceTypeEnumOutput values.
// You can construct a concrete instance of `ResourceTypeEnumInput` via:
//
//          ResourceTypeEnumArgs{...}
type ResourceTypeEnumInput interface {
	pulumi.Input

	ToResourceTypeEnumOutput() ResourceTypeEnumOutput
	ToResourceTypeEnumOutputWithContext(context.Context) ResourceTypeEnumOutput
}

var resourceTypeEnumPtrType = reflect.TypeOf((**ResourceTypeEnum)(nil)).Elem()

type ResourceTypeEnumPtrInput interface {
	pulumi.Input

	ToResourceTypeEnumPtrOutput() ResourceTypeEnumPtrOutput
	ToResourceTypeEnumPtrOutputWithContext(context.Context) ResourceTypeEnumPtrOutput
}

type resourceTypeEnumPtr string

func ResourceTypeEnumPtr(v string) ResourceTypeEnumPtrInput {
	return (*resourceTypeEnumPtr)(&v)
}

func (*resourceTypeEnumPtr) ElementType() reflect.Type {
	return resourceTypeEnumPtrType
}

func (in *resourceTypeEnumPtr) ToResourceTypeEnumPtrOutput() ResourceTypeEnumPtrOutput {
	return pulumi.ToOutput(in).(ResourceTypeEnumPtrOutput)
}

func (in *resourceTypeEnumPtr) ToResourceTypeEnumPtrOutputWithContext(ctx context.Context) ResourceTypeEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceTypeEnumPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ExampleEnumOutput{})
	pulumi.RegisterOutputType(ExampleEnumPtrOutput{})
	pulumi.RegisterOutputType(ExampleEnumInputEnumOutput{})
	pulumi.RegisterOutputType(ExampleEnumInputEnumPtrOutput{})
	pulumi.RegisterOutputType(ResourceTypeEnumOutput{})
	pulumi.RegisterOutputType(ResourceTypeEnumPtrOutput{})
}
