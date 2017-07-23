package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssert_default_values(t *testing.T) {
	root := YGNodeNew()

	assert.Equal(t, 0, YGNodeGetChildCount(root))
	var nilNode *YGNode
	assert.Equal(t, nilNode, YGNodeGetChild(root, 1))
	assert.Equal(t, nilNode, YGNodeGetChild(root, 0))

	assert.Equal(t, DirectionInherit, YGNodeStyleGetDirection(root))
	assert.Equal(t, FlexDirectionColumn, YGNodeStyleGetFlexDirection(root))
	assert.Equal(t, JustifyFlexStart, YGNodeStyleGetJustifyContent(root))
	assert.Equal(t, AlignFlexStart, YGNodeStyleGetAlignContent(root))
	assert.Equal(t, AlignStretch, YGNodeStyleGetAlignItems(root))
	assert.Equal(t, AlignAuto, YGNodeStyleGetAlignSelf(root))
	assert.Equal(t, PositionTypeRelative, YGNodeStyleGetPositionType(root))
	assert.Equal(t, YGWrapNoWrap, YGNodeStyleGetFlexWrap(root))
	assert.Equal(t, OverflowVisible, YGNodeStyleGetOverflow(root))
	assertFloatEqual(t, 0, YGNodeStyleGetFlexGrow(root))
	assertFloatEqual(t, 0, YGNodeStyleGetFlexShrink(root))
	assert.Equal(t, YGNodeStyleGetFlexBasis(root).unit, UnitAuto)

	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeLeft).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeTop).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeRight).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeBottom).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeStart).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPosition(root, EdgeEnd).unit, UnitUndefined)

	assert.Equal(t, YGNodeStyleGetMargin(root, EdgeLeft).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetMargin(root, EdgeTop).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetMargin(root, EdgeRight).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetMargin(root, EdgeBottom).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetMargin(root, EdgeStart).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetMargin(root, EdgeEnd).unit, UnitUndefined)

	assert.Equal(t, YGNodeStyleGetPadding(root, EdgeLeft).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPadding(root, EdgeTop).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPadding(root, EdgeRight).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPadding(root, EdgeBottom).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPadding(root, EdgeStart).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetPadding(root, EdgeEnd).unit, UnitUndefined)

	assert.True(t, YGFloatIsUndefined(YGNodeStyleGetBorder(root, EdgeLeft)))
	assert.True(t, YGFloatIsUndefined(YGNodeStyleGetBorder(root, EdgeTop)))
	assert.True(t, YGFloatIsUndefined(YGNodeStyleGetBorder(root, EdgeRight)))
	assert.True(t, YGFloatIsUndefined(YGNodeStyleGetBorder(root, EdgeBottom)))
	assert.True(t, YGFloatIsUndefined(YGNodeStyleGetBorder(root, EdgeStart)))
	assert.True(t, YGFloatIsUndefined(YGNodeStyleGetBorder(root, EdgeEnd)))

	assert.Equal(t, YGNodeStyleGetWidth(root).unit, UnitAuto)
	assert.Equal(t, YGNodeStyleGetHeight(root).unit, UnitAuto)
	assert.Equal(t, YGNodeStyleGetMinWidth(root).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetMinHeight(root).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetMaxWidth(root).unit, UnitUndefined)
	assert.Equal(t, YGNodeStyleGetMaxHeight(root).unit, UnitUndefined)

	assertFloatEqual(t, 0, YGNodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetTop(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetRight(root))
	assertFloatEqual(t, 0, YGNodeLayoutGetBottom(root))

	assertFloatEqual(t, 0, YGNodeLayoutGetMargin(root, EdgeLeft))
	assertFloatEqual(t, 0, YGNodeLayoutGetMargin(root, EdgeTop))
	assertFloatEqual(t, 0, YGNodeLayoutGetMargin(root, EdgeRight))
	assertFloatEqual(t, 0, YGNodeLayoutGetMargin(root, EdgeBottom))

	assertFloatEqual(t, 0, YGNodeLayoutGetPadding(root, EdgeLeft))
	assertFloatEqual(t, 0, YGNodeLayoutGetPadding(root, EdgeTop))
	assertFloatEqual(t, 0, YGNodeLayoutGetPadding(root, EdgeRight))
	assertFloatEqual(t, 0, YGNodeLayoutGetPadding(root, EdgeBottom))

	assertFloatEqual(t, 0, YGNodeLayoutGetBorder(root, EdgeLeft))
	assertFloatEqual(t, 0, YGNodeLayoutGetBorder(root, EdgeTop))
	assertFloatEqual(t, 0, YGNodeLayoutGetBorder(root, EdgeRight))
	assertFloatEqual(t, 0, YGNodeLayoutGetBorder(root, EdgeBottom))

	assert.True(t, YGFloatIsUndefined(YGNodeLayoutGetWidth(root)))
	assert.True(t, YGFloatIsUndefined(YGNodeLayoutGetHeight(root)))
	assert.Equal(t, DirectionInherit, YGNodeLayoutGetDirection(root))

}

func TestAssert_webdefault_values(t *testing.T) {
	config := YGConfigNew()
	YGConfigSetUseWebDefaults(config, true)
	root := YGNodeNewWithConfig(config)

	assert.Equal(t, FlexDirectionRow, YGNodeStyleGetFlexDirection(root))
	assert.Equal(t, AlignStretch, YGNodeStyleGetAlignContent(root))
	assertFloatEqual(t, 1, YGNodeStyleGetFlexShrink(root))

}

func TestAssert_webdefault_values_reset(t *testing.T) {
	config := YGConfigNew()
	YGConfigSetUseWebDefaults(config, true)
	root := YGNodeNewWithConfig(config)
	YGNodeReset(root)

	assert.Equal(t, FlexDirectionRow, YGNodeStyleGetFlexDirection(root))
	assert.Equal(t, AlignStretch, YGNodeStyleGetAlignContent(root))
	assertFloatEqual(t, 1, YGNodeStyleGetFlexShrink(root))

}
