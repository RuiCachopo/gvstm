package locks

import "gvstm/stmbench7/interfaces"

var (
    LocksInitialiser = interfaces.SynchMethodInitialiser{
        DOFactory: doFactory,
        BEFactory: beFactory,
    }
    doFactory = interfaces.DesignObjFactory{
        CreateAtomicPart:      newAtomicPartImpl,
        CreateConnection:      interfaces.NewConnectionImpl,
        CreateBaseAssembly:    newBaseAssemblyImpl,
        CreateComplexAssembly: newComplexAssemblyImpl,
        CreateCompositePart:   newCompositePartImpl,
        CreateDocument:        newDocumentImpl,
        CreateManual:          newManualImpl,
        CreateModule:          newModuleImpl,
    }
    beFactory = interfaces.BackendFactory{
        CreateLargeSet: newLargeSetImpl,
        CreateIndex:    newIndexImpl,
        CreateIDPool:   newIDPoolImpl,
    }
)