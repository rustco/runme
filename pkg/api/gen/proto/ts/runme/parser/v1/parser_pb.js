/* eslint-disable */
// @generated by protobuf-ts 2.9.4 with parameter output_javascript,optimize_code_size,long_type_string,add_pb_suffix,ts_nocheck,eslint_disable
// @generated from protobuf file "runme/parser/v1/parser.proto" (package "runme.parser.v1", syntax proto3)
// tslint:disable
// @ts-nocheck
/* eslint-disable */
// @generated by protobuf-ts 2.9.4 with parameter output_javascript,optimize_code_size,long_type_string,add_pb_suffix,ts_nocheck,eslint_disable
// @generated from protobuf file "runme/parser/v1/parser.proto" (package "runme.parser.v1", syntax proto3)
// tslint:disable
// @ts-nocheck
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { MessageType } from "@protobuf-ts/runtime";
import { BoolValue } from "../../../google/protobuf/wrappers_pb";
import { UInt32Value } from "../../../google/protobuf/wrappers_pb";
import { Int64Value } from "../../../google/protobuf/wrappers_pb";
/**
 * @generated from protobuf enum runme.parser.v1.CellKind
 */
export var CellKind;
(function (CellKind) {
    /**
     * @generated from protobuf enum value: CELL_KIND_UNSPECIFIED = 0;
     */
    CellKind[CellKind["UNSPECIFIED"] = 0] = "UNSPECIFIED";
    /**
     * @generated from protobuf enum value: CELL_KIND_MARKUP = 1;
     */
    CellKind[CellKind["MARKUP"] = 1] = "MARKUP";
    /**
     * @generated from protobuf enum value: CELL_KIND_CODE = 2;
     */
    CellKind[CellKind["CODE"] = 2] = "CODE";
})(CellKind || (CellKind = {}));
/**
 * @generated from protobuf enum runme.parser.v1.RunmeIdentity
 */
export var RunmeIdentity;
(function (RunmeIdentity) {
    /**
     * aka NONE
     *
     * @generated from protobuf enum value: RUNME_IDENTITY_UNSPECIFIED = 0;
     */
    RunmeIdentity[RunmeIdentity["UNSPECIFIED"] = 0] = "UNSPECIFIED";
    /**
     * @generated from protobuf enum value: RUNME_IDENTITY_ALL = 1;
     */
    RunmeIdentity[RunmeIdentity["ALL"] = 1] = "ALL";
    /**
     * @generated from protobuf enum value: RUNME_IDENTITY_DOCUMENT = 2;
     */
    RunmeIdentity[RunmeIdentity["DOCUMENT"] = 2] = "DOCUMENT";
    /**
     * @generated from protobuf enum value: RUNME_IDENTITY_CELL = 3;
     */
    RunmeIdentity[RunmeIdentity["CELL"] = 3] = "CELL";
})(RunmeIdentity || (RunmeIdentity = {}));
// @generated message type with reflection information, may provide speed optimized methods
class Notebook$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.Notebook", [
            { no: 1, name: "cells", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Cell },
            { no: 2, name: "metadata", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } },
            { no: 3, name: "frontmatter", kind: "message", T: () => Frontmatter }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.Notebook
 */
export const Notebook = new Notebook$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ExecutionSummaryTiming$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.ExecutionSummaryTiming", [
            { no: 1, name: "start_time", kind: "message", T: () => Int64Value },
            { no: 2, name: "end_time", kind: "message", T: () => Int64Value }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.ExecutionSummaryTiming
 */
export const ExecutionSummaryTiming = new ExecutionSummaryTiming$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CellOutputItem$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.CellOutputItem", [
            { no: 1, name: "data", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 2, name: "type", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "mime", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.CellOutputItem
 */
export const CellOutputItem = new CellOutputItem$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ProcessInfoExitReason$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.ProcessInfoExitReason", [
            { no: 1, name: "type", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "code", kind: "message", T: () => UInt32Value }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.ProcessInfoExitReason
 */
export const ProcessInfoExitReason = new ProcessInfoExitReason$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CellOutputProcessInfo$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.CellOutputProcessInfo", [
            { no: 1, name: "exit_reason", kind: "message", T: () => ProcessInfoExitReason },
            { no: 2, name: "pid", kind: "message", T: () => Int64Value }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.CellOutputProcessInfo
 */
export const CellOutputProcessInfo = new CellOutputProcessInfo$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CellOutput$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.CellOutput", [
            { no: 1, name: "items", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => CellOutputItem },
            { no: 2, name: "metadata", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } },
            { no: 3, name: "process_info", kind: "message", T: () => CellOutputProcessInfo }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.CellOutput
 */
export const CellOutput = new CellOutput$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CellExecutionSummary$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.CellExecutionSummary", [
            { no: 1, name: "execution_order", kind: "message", T: () => UInt32Value },
            { no: 2, name: "success", kind: "message", T: () => BoolValue },
            { no: 3, name: "timing", kind: "message", T: () => ExecutionSummaryTiming }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.CellExecutionSummary
 */
export const CellExecutionSummary = new CellExecutionSummary$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TextRange$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.TextRange", [
            { no: 1, name: "start", kind: "scalar", T: 13 /*ScalarType.UINT32*/ },
            { no: 2, name: "end", kind: "scalar", T: 13 /*ScalarType.UINT32*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.TextRange
 */
export const TextRange = new TextRange$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Cell$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.Cell", [
            { no: 1, name: "kind", kind: "enum", T: () => ["runme.parser.v1.CellKind", CellKind, "CELL_KIND_"] },
            { no: 2, name: "value", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "language_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "metadata", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } },
            { no: 5, name: "text_range", kind: "message", T: () => TextRange },
            { no: 6, name: "outputs", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => CellOutput },
            { no: 7, name: "execution_summary", kind: "message", T: () => CellExecutionSummary }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.Cell
 */
export const Cell = new Cell$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RunmeSessionDocument$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.RunmeSessionDocument", [
            { no: 1, name: "relative_path", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.RunmeSessionDocument
 */
export const RunmeSessionDocument = new RunmeSessionDocument$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RunmeSession$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.RunmeSession", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "document", kind: "message", T: () => RunmeSessionDocument }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.RunmeSession
 */
export const RunmeSession = new RunmeSession$Type();
// @generated message type with reflection information, may provide speed optimized methods
class FrontmatterRunme$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.FrontmatterRunme", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "version", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "session", kind: "message", T: () => RunmeSession }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.FrontmatterRunme
 */
export const FrontmatterRunme = new FrontmatterRunme$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Frontmatter$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.Frontmatter", [
            { no: 1, name: "shell", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "cwd", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "skip_prompts", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "runme", kind: "message", T: () => FrontmatterRunme },
            { no: 5, name: "category", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 6, name: "terminal_rows", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.Frontmatter
 */
export const Frontmatter = new Frontmatter$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeserializeRequestOptions$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.DeserializeRequestOptions", [
            { no: 1, name: "identity", kind: "enum", T: () => ["runme.parser.v1.RunmeIdentity", RunmeIdentity, "RUNME_IDENTITY_"] }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.DeserializeRequestOptions
 */
export const DeserializeRequestOptions = new DeserializeRequestOptions$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeserializeRequest$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.DeserializeRequest", [
            { no: 1, name: "source", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 2, name: "options", kind: "message", T: () => DeserializeRequestOptions }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.DeserializeRequest
 */
export const DeserializeRequest = new DeserializeRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeserializeResponse$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.DeserializeResponse", [
            { no: 1, name: "notebook", kind: "message", T: () => Notebook }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.DeserializeResponse
 */
export const DeserializeResponse = new DeserializeResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SerializeRequestOutputOptions$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.SerializeRequestOutputOptions", [
            { no: 1, name: "enabled", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "summary", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.SerializeRequestOutputOptions
 */
export const SerializeRequestOutputOptions = new SerializeRequestOutputOptions$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SerializeRequestOptions$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.SerializeRequestOptions", [
            { no: 1, name: "outputs", kind: "message", T: () => SerializeRequestOutputOptions },
            { no: 2, name: "session", kind: "message", T: () => RunmeSession }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.SerializeRequestOptions
 */
export const SerializeRequestOptions = new SerializeRequestOptions$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SerializeRequest$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.SerializeRequest", [
            { no: 1, name: "notebook", kind: "message", T: () => Notebook },
            { no: 2, name: "options", kind: "message", T: () => SerializeRequestOptions }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.SerializeRequest
 */
export const SerializeRequest = new SerializeRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SerializeResponse$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.SerializeResponse", [
            { no: 1, name: "result", kind: "scalar", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.SerializeResponse
 */
export const SerializeResponse = new SerializeResponse$Type();
/**
 * @generated ServiceType for protobuf service runme.parser.v1.ParserService
 */
export const ParserService = new ServiceType("runme.parser.v1.ParserService", [
    { name: "Deserialize", options: {}, I: DeserializeRequest, O: DeserializeResponse },
    { name: "Serialize", options: {}, I: SerializeRequest, O: SerializeResponse }
]);
