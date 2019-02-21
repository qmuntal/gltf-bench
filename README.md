# gltf-bench
Benchmarks and tests for gltf qmuntal/gltf package using the official [glTF Samples](https://github.com/KhronosGroup/glTF-Sample-Models).

## Benchmarks for v0.6.4
goos: windows
goarch: amd64
### ASCII
```
2CylinderEngine.gltf-8               300	   5501156 ns/op	 2431753 B/op	    4844 allocs/op
AlphaBlendModeTest.gltf-8           1000	   1678540 ns/op	  248817 B/op	    1581 allocs/op
AnimatedCube.gltf-8                 3000	    532586 ns/op	   42656 B/op	     315 allocs/op
AnimatedMorphCube.gltf-8            3000	    504992 ns/op	   49184 B/op	     366 allocs/op
AnimatedMorphSphere.gltf-8          3000	    567003 ns/op	  265568 B/op	     360 allocs/op
AnimatedTriangle.gltf-8             3000	    514298 ns/op	   24936 B/op	     185 allocs/op
AntiqueCamera.gltf-8                2000	    900129 ns/op	  864692 B/op	     394 allocs/op
Avocado.gltf-8                      5000	    421985 ns/op	   51184 B/op	     210 allocs/op
BarramundiFish.gltf-8               3000	    451322 ns/op	  158288 B/op	     214 allocs/op
BoomBox.gltf-8                      3000	    492362 ns/op	  240400 B/op	     225 allocs/op
BoomBoxWithAxes.gltf-8              2000	   1023764 ns/op	  687425 B/op	     715 allocs/op
Box.gltf-8                          5000	    406512 ns/op	   23056 B/op	     181 allocs/op
BoxAnimated.gltf-8                  2000	    680697 ns/op	   54224 B/op	     381 allocs/op
BoxInterleaved.gltf-8               5000	    432043 ns/op	   23792 B/op	     183 allocs/op
BoxTextured.gltf-8                  3000	    462430 ns/op	   34256 B/op	     228 allocs/op
BoxTexturedNonPowerOfTwo.gltf-8     3000	    455116 ns/op	   34624 B/op	     228 allocs/op
BoxVertexColors.gltf-8              3000	    528597 ns/op	   36480 B/op	     253 allocs/op
BrainStem.gltf-8                     100	  11449575 ns/op	 4294497 B/op	   10398 allocs/op
Buggy.gltf-8                         100	  21352919 ns/op	 9966818 B/op	   17606 allocs/op
Cameras.gltf-8                      5000	    350071 ns/op	   22024 B/op	     150 allocs/op
CesiumMan.gltf-8                     500	   3414863 ns/op	  734625 B/op	    2959 allocs/op
CesiumMilkTruck.gltf-8              2000	   1032238 ns/op	  191456 B/op	     659 allocs/op
Corset.gltf-8                       2000	    602253 ns/op	  690451 B/op	     214 allocs/op
Cube.gltf-8                         3000	    449129 ns/op	   37264 B/op	     245 allocs/op
DamagedHelmet.gltf-8                2000	    743023 ns/op	  601712 B/op	     251 allocs/op
Duck.gltf-8                         3000	    541737 ns/op	  141616 B/op	     259 allocs/op
EnvironmentTest.gltf-8              2000	    799414 ns/op	  407104 B/op	     402 allocs/op
FlightHelmet.gltf-8                  500	   2355702 ns/op	 3577776 B/op	     853 allocs/op
GearboxAssy.gltf-8                   200	   7415176 ns/op	 5626864 B/op	    5764 allocs/op
InterpolationTest.gltf-8            1000	   1945860 ns/op	  276561 B/op	    1608 allocs/op
Lantern.gltf-8                      2000	    726565 ns/op	  290673 B/op	     452 allocs/op
MetalRoughSpheres.gltf-8             300	   5947545 ns/op	11285552 B/op	     621 allocs/op
Monster.gltf-8                       300	   5349159 ns/op	  817888 B/op	    4763 allocs/op
MorphPrimitivesTest.gltf-8          3000	    587097 ns/op	   47184 B/op	     373 allocs/op
MultiUVTest.gltf-8                  3000	    579786 ns/op	   39744 B/op	     309 allocs/op
NormalTangentMirrorTest.gltf-8      3000	    545883 ns/op	  209441 B/op	     261 allocs/op
NormalTangentTest.gltf-8            3000	    576477 ns/op	  215760 B/op	     234 allocs/op
OrientationTest.gltf-8              1000	   1824149 ns/op	  261840 B/op	    1457 allocs/op
ReciprocatingSaw.gltf-8              100	  10322264 ns/op	 5091809 B/op	    9390 allocs/op
RiggedFigure.gltf-8                  500	   3325085 ns/op	  503185 B/op	    2842 allocs/op
RiggedSimple.gltf-8                 2000	    899662 ns/op	   81568 B/op	     593 allocs/op
SciFiHelmet.gltf-8                  1000	   1966739 ns/op	 3682752 B/op	     275 allocs/op
SimpleMeshes.gltf-8                 5000	    319950 ns/op	   17512 B/op	     143 allocs/op
SimpleMorph.gltf-8                  3000	    497680 ns/op	   27528 B/op	     235 allocs/op
SimpleSparseAccessor.gltf-8         5000	    324540 ns/op	   21672 B/op	     131 allocs/op
SpecGlossVsMetalRough.gltf-8        2000	    970078 ns/op	  233896 B/op	     615 allocs/op
Sponza.gltf-8                        100	  13265795 ns/op	10808791 B/op	   10416 allocs/op
Suzanne.gltf-8                      2000	    595479 ns/op	  633536 B/op	     247 allocs/op
TextureCoordinateTest.gltf-8        2000	    932990 ns/op	   85728 B/op	     721 allocs/op
TextureSettingsTest.gltf-8          1000	   1427221 ns/op	  167312 B/op	    1296 allocs/op
TextureTransformTest.gltf-8         2000	    765951 ns/op	   92896 B/op	     692 allocs/op
Triangle.gltf-8                     5000	    286042 ns/op	   16384 B/op	     114 allocs/op
TriangleWithoutIndices.gltf-8       5000	    274875 ns/op	   16136 B/op	     100 allocs/op
TwoSidedPlane.gltf-8                3000	    453773 ns/op	   36528 B/op	     249 allocs/op
UnlitTest.gltf-8                    3000	    434185 ns/op	   36576 B/op	     195 allocs/op
VC.gltf-8                             50	  27446576 ns/op	 6053961 B/op	   27044 allocs/op
VertexColorTest.gltf-8              2000	    723048 ns/op	   68752 B/op	     455 allocs/op
WaterBottle.gltf-8                  3000	    471404 ns/op	  183168 B/op	     225 allocs/op
```
### Embedded
```
2CylinderEngine.gltf-8                30	  53527030 ns/op	15357528 B/op	    4841 allocs/op
AlphaBlendModeTest.gltf-8             20	  81880115 ns/op	12529835 B/op	    1579 allocs/op
AnimatedTriangle.gltf-8             5000	    266893 ns/op	   23064 B/op	     169 allocs/op
Box.gltf-8                          5000	    298807 ns/op	   32568 B/op	     178 allocs/op
BoxAnimated.gltf-8                  2000	    777917 ns/op	  129464 B/op	     375 allocs/op
BoxInterleaved.gltf-8               5000	    295018 ns/op	   32728 B/op	     176 allocs/op
BoxTextured.gltf-8                  2000	    971899 ns/op	  183064 B/op	     223 allocs/op
BoxTexturedNonPowerOfTwo.gltf-8     3000	    470741 ns/op	   55672 B/op	     221 allocs/op
BoxVertexColors.gltf-8              3000	    399605 ns/op	   38968 B/op	     245 allocs/op
BrainStem.gltf-8                      20	  96294090 ns/op	28836680 B/op	   10395 allocs/op
Buggy.gltf-8                           5	 238170360 ns/op	63131864 B/op	   17603 allocs/op
Cameras.gltf-8                      5000	    243356 ns/op	   21176 B/op	     142 allocs/op
CesiumMan.gltf-8                     100	  15628725 ns/op	 3543420 B/op	    2954 allocs/op
CesiumMilkTruck.gltf-8               100	  14152100 ns/op	 3122968 B/op	     657 allocs/op
DamagedHelmet.gltf-8                  20	  97289385 ns/op	23153000 B/op	     253 allocs/op
Duck.gltf-8                          500	   3701342 ns/op	  948809 B/op	     256 allocs/op
GearboxAssy.gltf-8                    10	 153888550 ns/op	35264470 B/op	    5762 allocs/op
MetalRoughSpheres.gltf-8               5	 289231020 ns/op	74698568 B/op	     623 allocs/op
Monster.gltf-8                       100	  12875372 ns/op	 2250456 B/op	    4757 allocs/op
MultiUVTest.gltf-8                  1000	   1504014 ns/op	  210696 B/op	     304 allocs/op
NormalTangentMirrorTest.gltf-8        30	  52726780 ns/op	11496552 B/op	     262 allocs/op
NormalTangentTest.gltf-8              30	  48171130 ns/op	11396557 B/op	     235 allocs/op
OrientationTest.gltf-8              1000	   2506293 ns/op	  473640 B/op	    1450 allocs/op
ReciprocatingSaw.gltf-8               10	 105019260 ns/op	30682424 B/op	    9387 allocs/op
RiggedFigure.gltf-8                  300	   3909508 ns/op	  567592 B/op	    2834 allocs/op
RiggedSimple.gltf-8                 2000	    998847 ns/op	  130856 B/op	     586 allocs/op
SimpleMeshes.gltf-8                 5000	    236566 ns/op	   20840 B/op	     136 allocs/op
TextureCoordinateTest.gltf-8        2000	   1036745 ns/op	  129288 B/op	     714 allocs/op
TextureSettingsTest.gltf-8          1000	   2347720 ns/op	  413576 B/op	    1290 allocs/op
Triangle.gltf-8                    10000	    190295 ns/op	   15488 B/op	     106 allocs/op
TriangleWithoutIndices.gltf-8      10000	    183911 ns/op	   14952 B/op	      92 allocs/op
VC.gltf-8                             10	 100633010 ns/op	27221768 B/op	   27043 allocs/op
VertexColorTest.gltf-8              2000	   1116087 ns/op	  203944 B/op	     449 allocs/op
```
### Binary
```
2CylinderEngine.glb-8                500	   3444836 ns/op	 2299548 B/op	    4838 allocs/op
AlphaBlendModeTest.glb-8            1000	   1701541 ns/op	 3149278 B/op	    1448 allocs/op
AnimatedMorphCube.glb-8             5000	    310975 ns/op	   39848 B/op	     360 allocs/op
AnimatedMorphSphere.glb-8           5000	    359515 ns/op	  256200 B/op	     354 allocs/op
AntiqueCamera.glb-8                 1000	   2323110 ns/op	19003760 B/op	     426 allocs/op
Avocado.glb-8                       1000	   1237908 ns/op	 8353985 B/op	     221 allocs/op
BarramundiFish.glb-8                1000	   1604923 ns/op	12552799 B/op	     226 allocs/op
BoomBox.glb-8                       1000	   1476170 ns/op	10972302 B/op	     242 allocs/op
Box.glb-8                          10000	    197867 ns/op	   18424 B/op	     179 allocs/op
BoxAnimated.glb-8                   5000	    332119 ns/op	   45048 B/op	     375 allocs/op
BoxInterleaved.glb-8               10000	    197869 ns/op	   18568 B/op	     177 allocs/op
BoxTextured.glb-8                  10000	    237358 ns/op	   44824 B/op	     204 allocs/op
BoxTexturedNonPowerOfTwo.glb-8      5000	    252125 ns/op	   23912 B/op	     226 allocs/op
BoxVertexColors.glb-8               5000	    283846 ns/op	   27112 B/op	     206 allocs/op
BrainStem.glb-8                      200	   7056277 ns/op	 4031455 B/op	   10392 allocs/op
Buggy.glb-8                          100	  10262534 ns/op	 9441650 B/op	   17600 allocs/op
CesiumMan.glb-8                     1000	   2197097 ns/op	  750314 B/op	    2957 allocs/op
CesiumMilkTruck.glb-8               2000	    646366 ns/op	  600138 B/op	     658 allocs/op
Corset.glb-8                        1000	   2137780 ns/op	13609404 B/op	     226 allocs/op
DamagedHelmet.glb-8                 2000	    983181 ns/op	 3805550 B/op	     266 allocs/op
Duck.glb-8                          5000	    343288 ns/op	  149080 B/op	     258 allocs/op
GearboxAssy.glb-8                    500	   3911601 ns/op	 5494837 B/op	    5758 allocs/op
InterpolationTest.glb-8             2000	   1062671 ns/op	  188392 B/op	    1606 allocs/op
Lantern.glb-8                       1000	   1608854 ns/op	 9924400 B/op	     471 allocs/op
MetalRoughSpheres.glb-8             1000	   1729625 ns/op	11284997 B/op	     594 allocs/op
Monster.glb-8                        500	   2984239 ns/op	  776233 B/op	    4762 allocs/op
MorphPrimitivesTest.glb-8           5000	    336116 ns/op	   93880 B/op	     387 allocs/op
MultiUVTest.glb-8                   5000	    366220 ns/op	   71048 B/op	     315 allocs/op
NormalTangentMirrorTest.glb-8       2000	    745759 ns/op	 2052446 B/op	     269 allocs/op
NormalTangentTest.glb-8             2000	    680282 ns/op	 1952272 B/op	     241 allocs/op
OrientationTest.glb-8               2000	    915574 ns/op	  162088 B/op	    1387 allocs/op
ReciprocatingSaw.glb-8               300	   5266690 ns/op	 4828532 B/op	    9384 allocs/op
RiggedFigure.glb-8                  1000	   1884958 ns/op	  305496 B/op	    2835 allocs/op
RiggedSimple.glb-8                  3000	    496682 ns/op	   64072 B/op	     587 allocs/op
SpecGlossVsMetalRough.glb-8         1000	   2320692 ns/op	15529408 B/op	     625 allocs/op
TextureCoordinateTest.glb-8         3000	    562520 ns/op	   77080 B/op	     695 allocs/op
TextureSettingsTest.glb-8           2000	    884155 ns/op	  161400 B/op	    1222 allocs/op
UnlitTest.glb-8                    10000	    224197 ns/op	   23320 B/op	     188 allocs/op
VC.glb-8                             100	  15299429 ns/op	 5365575 B/op	   27165 allocs/op
VertexColorTest.glb-8               5000	    392158 ns/op	   63048 B/op	     424 allocs/op
WaterBottle.glb-8                   1000	   1347806 ns/op	 8989918 B/op	     242 allocs/op
```