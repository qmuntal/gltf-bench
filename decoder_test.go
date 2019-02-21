package gltfbench

import (
	"testing"

	. "github.com/qmuntal/gltf"
)

func TestOpenASCII(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"testdata/2.0/2CylinderEngine/glTF/2CylinderEngine.gltf"},
		{"testdata/2.0/AlphaBlendModeTest/glTF/AlphaBlendModeTest.gltf"},
		{"testdata/2.0/AnimatedCube/glTF/AnimatedCube.gltf"},
		{"testdata/2.0/AnimatedMorphCube/glTF/AnimatedMorphCube.gltf"},
		{"testdata/2.0/AnimatedMorphSphere/glTF/AnimatedMorphSphere.gltf"},
		{"testdata/2.0/AnimatedTriangle/glTF/AnimatedTriangle.gltf"},
		{"testdata/2.0/AntiqueCamera/glTF/AntiqueCamera.gltf"},
		{"testdata/2.0/Avocado/glTF/Avocado.gltf"},
		{"testdata/2.0/BarramundiFish/glTF/BarramundiFish.gltf"},
		{"testdata/2.0/BoomBox/glTF/BoomBox.gltf"},
		{"testdata/2.0/BoomBoxWithAxes/glTF/BoomBoxWithAxes.gltf"},
		{"testdata/2.0/Box/glTF/Box.gltf"},
		{"testdata/2.0/BoxAnimated/glTF/BoxAnimated.gltf"},
		{"testdata/2.0/BoxInterleaved/glTF/BoxInterleaved.gltf"},
		{"testdata/2.0/BoxTextured/glTF/BoxTextured.gltf"},
		{"testdata/2.0/BoxTexturedNonPowerOfTwo/glTF/BoxTexturedNonPowerOfTwo.gltf"},
		{"testdata/2.0/BoxVertexColors/glTF/BoxVertexColors.gltf"},
		{"testdata/2.0/BrainStem/glTF/BrainStem.gltf"},
		{"testdata/2.0/Buggy/glTF/Buggy.gltf"},
		{"testdata/2.0/Cameras/glTF/Cameras.gltf"},
		{"testdata/2.0/CesiumMan/glTF/CesiumMan.gltf"},
		{"testdata/2.0/CesiumMilkTruck/glTF/CesiumMilkTruck.gltf"},
		{"testdata/2.0/Corset/glTF/Corset.gltf"},
		{"testdata/2.0/Cube/glTF/Cube.gltf"},
		{"testdata/2.0/DamagedHelmet/glTF/DamagedHelmet.gltf"},
		{"testdata/2.0/Duck/glTF/Duck.gltf"},
		{"testdata/2.0/EnvironmentTest/glTF/EnvironmentTest.gltf"},
		{"testdata/2.0/FlightHelmet/glTF/FlightHelmet.gltf"},
		{"testdata/2.0/GearboxAssy/glTF/GearboxAssy.gltf"},
		{"testdata/2.0/InterpolationTest/glTF/InterpolationTest.gltf"},
		{"testdata/2.0/Lantern/glTF/Lantern.gltf"},
		{"testdata/2.0/MetalRoughSpheres/glTF/MetalRoughSpheres.gltf"},
		{"testdata/2.0/Monster/glTF/Monster.gltf"},
		{"testdata/2.0/MorphPrimitivesTest/glTF/MorphPrimitivesTest.gltf"},
		{"testdata/2.0/MultiUVTest/glTF/MultiUVTest.gltf"},
		{"testdata/2.0/NormalTangentMirrorTest/glTF/NormalTangentMirrorTest.gltf"},
		{"testdata/2.0/NormalTangentTest/glTF/NormalTangentTest.gltf"},
		{"testdata/2.0/OrientationTest/glTF/OrientationTest.gltf"},
		{"testdata/2.0/ReciprocatingSaw/glTF/ReciprocatingSaw.gltf"},
		{"testdata/2.0/RiggedFigure/glTF/RiggedFigure.gltf"},
		{"testdata/2.0/RiggedSimple/glTF/RiggedSimple.gltf"},
		{"testdata/2.0/SciFiHelmet/glTF/SciFiHelmet.gltf"},
		{"testdata/2.0/SimpleMeshes/glTF/SimpleMeshes.gltf"},
		{"testdata/2.0/SimpleMorph/glTF/SimpleMorph.gltf"},
		{"testdata/2.0/SimpleSparseAccessor/glTF/SimpleSparseAccessor.gltf"},
		{"testdata/2.0/SpecGlossVsMetalRough/glTF/SpecGlossVsMetalRough.gltf"},
		{"testdata/2.0/Sponza/glTF/Sponza.gltf"},
		{"testdata/2.0/Suzanne/glTF/Suzanne.gltf"},
		{"testdata/2.0/TextureCoordinateTest/glTF/TextureCoordinateTest.gltf"},
		{"testdata/2.0/TextureSettingsTest/glTF/TextureSettingsTest.gltf"},
		{"testdata/2.0/TextureTransformTest/glTF/TextureTransformTest.gltf"},
		{"testdata/2.0/Triangle/glTF/Triangle.gltf"},
		{"testdata/2.0/TriangleWithoutIndices/glTF/TriangleWithoutIndices.gltf"},
		{"testdata/2.0/TwoSidedPlane/glTF/TwoSidedPlane.gltf"},
		{"testdata/2.0/UnlitTest/glTF/UnlitTest.gltf"},
		{"testdata/2.0/VC/glTF/VC.gltf"},
		{"testdata/2.0/VertexColorTest/glTF/VertexColorTest.gltf"},
		{"testdata/2.0/WaterBottle/glTF/WaterBottle.gltf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := Open(tt.name)
			if err != nil {
				t.Errorf("Open() error = %v", err)
				return
			}
			if err = doc.Validate(); err != nil {
				t.Errorf("Document.Validate() error = %v", err)
				return
			}
		})
	}
}

func TestOpenEmbedded(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"testdata/2.0/2CylinderEngine/glTF-Embedded/2CylinderEngine.gltf"},
		{"testdata/2.0/AlphaBlendModeTest/glTF-Embedded/AlphaBlendModeTest.gltf"},
		{"testdata/2.0/AnimatedTriangle/glTF-Embedded/AnimatedTriangle.gltf"},
		{"testdata/2.0/Box/glTF-Embedded/Box.gltf"},
		{"testdata/2.0/BoxAnimated/glTF-Embedded/BoxAnimated.gltf"},
		{"testdata/2.0/BoxInterleaved/glTF-Embedded/BoxInterleaved.gltf"},
		{"testdata/2.0/BoxTextured/glTF-Embedded/BoxTextured.gltf"},
		{"testdata/2.0/BoxTexturedNonPowerOfTwo/glTF-Embedded/BoxTexturedNonPowerOfTwo.gltf"},
		{"testdata/2.0/BoxVertexColors/glTF-Embedded/BoxVertexColors.gltf"},
		{"testdata/2.0/BrainStem/glTF-Embedded/BrainStem.gltf"},
		{"testdata/2.0/Buggy/glTF-Embedded/Buggy.gltf"},
		{"testdata/2.0/Cameras/glTF-Embedded/Cameras.gltf"},
		{"testdata/2.0/CesiumMan/glTF-Embedded/CesiumMan.gltf"},
		{"testdata/2.0/CesiumMilkTruck/glTF-Embedded/CesiumMilkTruck.gltf"},
		{"testdata/2.0/DamagedHelmet/glTF-Embedded/DamagedHelmet.gltf"},
		{"testdata/2.0/Duck/glTF-Embedded/Duck.gltf"},
		{"testdata/2.0/GearboxAssy/glTF-Embedded/GearboxAssy.gltf"},
		{"testdata/2.0/MetalRoughSpheres/glTF-Embedded/MetalRoughSpheres.gltf"},
		{"testdata/2.0/Monster/glTF-Embedded/Monster.gltf"},
		{"testdata/2.0/MultiUVTest/glTF-Embedded/MultiUVTest.gltf"},
		{"testdata/2.0/NormalTangentMirrorTest/glTF-Embedded/NormalTangentMirrorTest.gltf"},
		{"testdata/2.0/NormalTangentTest/glTF-Embedded/NormalTangentTest.gltf"},
		{"testdata/2.0/OrientationTest/glTF-Embedded/OrientationTest.gltf"},
		{"testdata/2.0/ReciprocatingSaw/glTF-Embedded/ReciprocatingSaw.gltf"},
		{"testdata/2.0/RiggedFigure/glTF-Embedded/RiggedFigure.gltf"},
		{"testdata/2.0/RiggedSimple/glTF-Embedded/RiggedSimple.gltf"},
		{"testdata/2.0/SimpleMeshes/glTF-Embedded/SimpleMeshes.gltf"},
		{"testdata/2.0/TextureCoordinateTest/glTF-Embedded/TextureCoordinateTest.gltf"},
		{"testdata/2.0/TextureSettingsTest/glTF-Embedded/TextureSettingsTest.gltf"},
		{"testdata/2.0/Triangle/glTF-Embedded/Triangle.gltf"},
		{"testdata/2.0/TriangleWithoutIndices/glTF-Embedded/TriangleWithoutIndices.gltf"},
		{"testdata/2.0/VC/glTF-Embedded/VC.gltf"},
		{"testdata/2.0/VertexColorTest/glTF-Embedded/VertexColorTest.gltf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := Open(tt.name)
			if err != nil {
				t.Errorf("Open() error = %v", err)
				return
			}
			if err = doc.Validate(); err != nil {
				t.Errorf("Document.Validate() error = %v", err)
				return
			}
		})
	}
}

func TestOpenBinary(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"testdata/2.0/2CylinderEngine/glTF-Binary/2CylinderEngine.glb"},
		{"testdata/2.0/AlphaBlendModeTest/glTF-Binary/AlphaBlendModeTest.glb"},
		{"testdata/2.0/AnimatedMorphCube/glTF-Binary/AnimatedMorphCube.glb"},
		{"testdata/2.0/AnimatedMorphSphere/glTF-Binary/AnimatedMorphSphere.glb"},
		{"testdata/2.0/AntiqueCamera/glTF-Binary/AntiqueCamera.glb"},
		{"testdata/2.0/Avocado/glTF-Binary/Avocado.glb"},
		{"testdata/2.0/BarramundiFish/glTF-Binary/BarramundiFish.glb"},
		{"testdata/2.0/BoomBox/glTF-Binary/BoomBox.glb"},
		{"testdata/2.0/Box/glTF-Binary/Box.glb"},
		{"testdata/2.0/BoxAnimated/glTF-Binary/BoxAnimated.glb"},
		{"testdata/2.0/BoxInterleaved/glTF-Binary/BoxInterleaved.glb"},
		{"testdata/2.0/BoxTextured/glTF-Binary/BoxTextured.glb"},
		{"testdata/2.0/BoxTexturedNonPowerOfTwo/glTF-Binary/BoxTexturedNonPowerOfTwo.glb"},
		{"testdata/2.0/BoxVertexColors/glTF-Binary/BoxVertexColors.glb"},
		{"testdata/2.0/BrainStem/glTF-Binary/BrainStem.glb"},
		{"testdata/2.0/Buggy/glTF-Binary/Buggy.glb"},
		{"testdata/2.0/CesiumMan/glTF-Binary/CesiumMan.glb"},
		{"testdata/2.0/CesiumMilkTruck/glTF-Binary/CesiumMilkTruck.glb"},
		{"testdata/2.0/Corset/glTF-Binary/Corset.glb"},
		{"testdata/2.0/DamagedHelmet/glTF-Binary/DamagedHelmet.glb"},
		{"testdata/2.0/Duck/glTF-Binary/Duck.glb"},
		{"testdata/2.0/GearboxAssy/glTF-Binary/GearboxAssy.glb"},
		{"testdata/2.0/InterpolationTest/glTF-Binary/InterpolationTest.glb"},
		{"testdata/2.0/Lantern/glTF-Binary/Lantern.glb"},
		{"testdata/2.0/MetalRoughSpheres/glTF-Binary/MetalRoughSpheres.glb"},
		{"testdata/2.0/Monster/glTF-Binary/Monster.glb"},
		{"testdata/2.0/MorphPrimitivesTest/glTF-Binary/MorphPrimitivesTest.glb"},
		{"testdata/2.0/MultiUVTest/glTF-Binary/MultiUVTest.glb"},
		{"testdata/2.0/NormalTangentMirrorTest/glTF-Binary/NormalTangentMirrorTest.glb"},
		{"testdata/2.0/NormalTangentTest/glTF-Binary/NormalTangentTest.glb"},
		{"testdata/2.0/OrientationTest/glTF-Binary/OrientationTest.glb"},
		{"testdata/2.0/ReciprocatingSaw/glTF-Binary/ReciprocatingSaw.glb"},
		{"testdata/2.0/RiggedFigure/glTF-Binary/RiggedFigure.glb"},
		{"testdata/2.0/RiggedSimple/glTF-Binary/RiggedSimple.glb"},
		{"testdata/2.0/SpecGlossVsMetalRough/glTF-Binary/SpecGlossVsMetalRough.glb"},
		{"testdata/2.0/TextureCoordinateTest/glTF-Binary/TextureCoordinateTest.glb"},
		{"testdata/2.0/TextureSettingsTest/glTF-Binary/TextureSettingsTest.glb"},
		{"testdata/2.0/UnlitTest/glTF-Binary/UnlitTest.glb"},
		{"testdata/2.0/VC/glTF-Binary/VC.glb"},
		{"testdata/2.0/VertexColorTest/glTF-Binary/VertexColorTest.glb"},
		{"testdata/2.0/WaterBottle/glTF-Binary/WaterBottle.glb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := Open(tt.name)
			if err != nil {
				t.Errorf("Open() error = %v", err)
				return
			}
			if err = doc.Validate(); err != nil {
				t.Errorf("Document.Validate() error = %v", err)
				return
			}
		})
	}
}
