package internal

import "math"

var (
    NumAtomsPerCompPart = 200
    NumConnectionsPerAtomicPart = 6
    DocumentSize = 20000
    ManualSize = 1000000
    NumCompPartsPerModule = 500
    NumSubAssemblies  = 3
    NumAssemblyLevels = 7
    NumModules = 1

    MinModuleDate = 1000
    MaxModuleDate = 1999
    MinAssDate = 1000
    MaxAssDate = 1999
    MinAtomicPartDate = 1000
    MaxAtomicPartDate = 1999
    MinOldCompositePartDate = 0
    MaxOldCompositePartDate = 999
    MinYoungCompositePartDate = 2000
    MaxYoungCompositePartDate = 2999
    YoungCompositePartFraction = 10
    NumTypes = 10

    InitialTotalCompParts = NumModules * NumCompPartsPerModule
    InitialTotalBaseAssemblies = int(math.Pow(float64(NumSubAssemblies), float64(NumAssemblyLevels-1)))
    InitialTotalComplexAssemblies = (1 - InitialTotalBaseAssemblies) / (1 - NumSubAssemblies)

    MaxCompParts = int(1.05 * float64(InitialTotalCompParts))
    MaxAtomicParts = MaxCompParts * NumAtomsPerCompPart
    MaxBaseAssemblies = int(1.05 * float64(InitialTotalBaseAssemblies))
    MaxComplexAssemblies = int(1.05 * float64(InitialTotalComplexAssemblies))
)
