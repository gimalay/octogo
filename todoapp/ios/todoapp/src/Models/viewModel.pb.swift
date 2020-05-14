// DO NOT EDIT.
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: viewModel.proto
//
// For information on using the generated types, please see the documentation:
//   https://github.com/apple/swift-protobuf/

import Foundation
import SwiftProtobuf

// If the compiler emits an error on this type, it is because this file
// was generated by a version of the `protoc` Swift plug-in that is
// incompatible with the version of SwiftProtobuf to which you are linking.
// Please ensure that you are building against the same version of the API
// that was used to generate this file.
fileprivate struct _GeneratedWithProtocGenSwiftVersion: SwiftProtobuf.ProtobufAPIVersionCheck {
  struct _2: SwiftProtobuf.ProtobufAPIVersion_2 {}
  typealias Version = _2
}

enum LocationType: SwiftProtobuf.Enum {
  typealias RawValue = Int
  case unknown // = 0
  case home // = 1457
  case project // = 1571
  case task // = 1648
  case UNRECOGNIZED(Int)

  init() {
    self = .unknown
  }

  init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .unknown
    case 1457: self = .home
    case 1571: self = .project
    case 1648: self = .task
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  var rawValue: Int {
    switch self {
    case .unknown: return 0
    case .home: return 1457
    case .project: return 1571
    case .task: return 1648
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension LocationType: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  static var allCases: [LocationType] = [
    .unknown,
    .home,
    .project,
    .task,
  ]
}

#endif  // swift(>=4.2)

struct Location {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var unknownFields = SwiftProtobuf.UnknownStorage()

  struct Home {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    var unknownFields = SwiftProtobuf.UnknownStorage()

    init() {}
  }

  struct Project {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    var projectID: Data = SwiftProtobuf.Internal.emptyData

    var unknownFields = SwiftProtobuf.UnknownStorage()

    init() {}
  }

  struct Task {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    var taskID: Data = SwiftProtobuf.Internal.emptyData

    var unknownFields = SwiftProtobuf.UnknownStorage()

    init() {}
  }

  struct AddTask {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    var unknownFields = SwiftProtobuf.UnknownStorage()

    init() {}
  }

  init() {}
}

struct ViewModel {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var unknownFields = SwiftProtobuf.UnknownStorage()

  struct Project {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    var id: Data = SwiftProtobuf.Internal.emptyData

    var name: String = String()

    var tasks: [ViewModel.Project.Task] = []

    var unknownFields = SwiftProtobuf.UnknownStorage()

    struct Task {
      // SwiftProtobuf.Message conformance is added in an extension below. See the
      // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
      // methods supported on all messages.

      var id: Data = SwiftProtobuf.Internal.emptyData

      var name: String = String()

      var emoji: String = String()

      var unknownFields = SwiftProtobuf.UnknownStorage()

      init() {}
    }

    init() {}
  }

  struct Task {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    var id: Data = SwiftProtobuf.Internal.emptyData

    var name: String = String()

    var emoji: String = String()

    var unknownFields = SwiftProtobuf.UnknownStorage()

    init() {}
  }

  struct Home {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    var projects: [ViewModel.Home.Project] = []

    var unknownFields = SwiftProtobuf.UnknownStorage()

    struct Project {
      // SwiftProtobuf.Message conformance is added in an extension below. See the
      // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
      // methods supported on all messages.

      var id: Data = SwiftProtobuf.Internal.emptyData

      var name: String = String()

      var tasks: [ViewModel.Home.Project.Task] = []

      var unknownFields = SwiftProtobuf.UnknownStorage()

      struct Task {
        // SwiftProtobuf.Message conformance is added in an extension below. See the
        // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
        // methods supported on all messages.

        var id: Data = SwiftProtobuf.Internal.emptyData

        var name: String = String()

        var emoji: String = String()

        var unknownFields = SwiftProtobuf.UnknownStorage()

        init() {}
      }

      init() {}
    }

    init() {}
  }

  init() {}
}

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "viewModel"

extension LocationType: SwiftProtobuf._ProtoNameProviding {
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "Unknown"),
    1457: .same(proto: "Home"),
    1571: .same(proto: "Project"),
    1648: .same(proto: "Task"),
  ]
}

extension Location: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".Location"
  static let _protobuf_nameMap = SwiftProtobuf._NameMap()

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let _ = try decoder.nextFieldNumber() {
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Location, rhs: Location) -> Bool {
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Location.Home: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = Location.protoMessageName + ".Home"
  static let _protobuf_nameMap = SwiftProtobuf._NameMap()

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let _ = try decoder.nextFieldNumber() {
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Location.Home, rhs: Location.Home) -> Bool {
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Location.Project: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = Location.protoMessageName + ".Project"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    2697: .same(proto: "ProjectID"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      switch fieldNumber {
      case 2697: try decoder.decodeSingularBytesField(value: &self.projectID)
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.projectID.isEmpty {
      try visitor.visitSingularBytesField(value: self.projectID, fieldNumber: 2697)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Location.Project, rhs: Location.Project) -> Bool {
    if lhs.projectID != rhs.projectID {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Location.Task: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = Location.protoMessageName + ".Task"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1717: .same(proto: "TaskID"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      switch fieldNumber {
      case 1717: try decoder.decodeSingularBytesField(value: &self.taskID)
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.taskID.isEmpty {
      try visitor.visitSingularBytesField(value: self.taskID, fieldNumber: 1717)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Location.Task, rhs: Location.Task) -> Bool {
    if lhs.taskID != rhs.taskID {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Location.AddTask: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = Location.protoMessageName + ".AddTask"
  static let _protobuf_nameMap = SwiftProtobuf._NameMap()

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let _ = try decoder.nextFieldNumber() {
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Location.AddTask, rhs: Location.AddTask) -> Bool {
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension ViewModel: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".ViewModel"
  static let _protobuf_nameMap = SwiftProtobuf._NameMap()

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let _ = try decoder.nextFieldNumber() {
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: ViewModel, rhs: ViewModel) -> Bool {
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension ViewModel.Project: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = ViewModel.protoMessageName + ".Project"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    4947: .same(proto: "ID"),
    4032: .same(proto: "name"),
    8856: .same(proto: "tasks"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      switch fieldNumber {
      case 4032: try decoder.decodeSingularStringField(value: &self.name)
      case 4947: try decoder.decodeSingularBytesField(value: &self.id)
      case 8856: try decoder.decodeRepeatedMessageField(value: &self.tasks)
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.name.isEmpty {
      try visitor.visitSingularStringField(value: self.name, fieldNumber: 4032)
    }
    if !self.id.isEmpty {
      try visitor.visitSingularBytesField(value: self.id, fieldNumber: 4947)
    }
    if !self.tasks.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.tasks, fieldNumber: 8856)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: ViewModel.Project, rhs: ViewModel.Project) -> Bool {
    if lhs.id != rhs.id {return false}
    if lhs.name != rhs.name {return false}
    if lhs.tasks != rhs.tasks {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension ViewModel.Project.Task: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = ViewModel.Project.protoMessageName + ".Task"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    5946: .same(proto: "ID"),
    8336: .same(proto: "name"),
    9790: .same(proto: "emoji"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      switch fieldNumber {
      case 5946: try decoder.decodeSingularBytesField(value: &self.id)
      case 8336: try decoder.decodeSingularStringField(value: &self.name)
      case 9790: try decoder.decodeSingularStringField(value: &self.emoji)
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.id.isEmpty {
      try visitor.visitSingularBytesField(value: self.id, fieldNumber: 5946)
    }
    if !self.name.isEmpty {
      try visitor.visitSingularStringField(value: self.name, fieldNumber: 8336)
    }
    if !self.emoji.isEmpty {
      try visitor.visitSingularStringField(value: self.emoji, fieldNumber: 9790)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: ViewModel.Project.Task, rhs: ViewModel.Project.Task) -> Bool {
    if lhs.id != rhs.id {return false}
    if lhs.name != rhs.name {return false}
    if lhs.emoji != rhs.emoji {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension ViewModel.Task: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = ViewModel.protoMessageName + ".Task"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    4987: .same(proto: "ID"),
    4932: .same(proto: "name"),
    2651: .same(proto: "emoji"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      switch fieldNumber {
      case 2651: try decoder.decodeSingularStringField(value: &self.emoji)
      case 4932: try decoder.decodeSingularStringField(value: &self.name)
      case 4987: try decoder.decodeSingularBytesField(value: &self.id)
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.emoji.isEmpty {
      try visitor.visitSingularStringField(value: self.emoji, fieldNumber: 2651)
    }
    if !self.name.isEmpty {
      try visitor.visitSingularStringField(value: self.name, fieldNumber: 4932)
    }
    if !self.id.isEmpty {
      try visitor.visitSingularBytesField(value: self.id, fieldNumber: 4987)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: ViewModel.Task, rhs: ViewModel.Task) -> Bool {
    if lhs.id != rhs.id {return false}
    if lhs.name != rhs.name {return false}
    if lhs.emoji != rhs.emoji {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension ViewModel.Home: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = ViewModel.protoMessageName + ".Home"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    6910: .same(proto: "projects"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      switch fieldNumber {
      case 6910: try decoder.decodeRepeatedMessageField(value: &self.projects)
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.projects.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.projects, fieldNumber: 6910)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: ViewModel.Home, rhs: ViewModel.Home) -> Bool {
    if lhs.projects != rhs.projects {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension ViewModel.Home.Project: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = ViewModel.Home.protoMessageName + ".Project"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    4947: .same(proto: "ID"),
    4032: .same(proto: "name"),
    8856: .same(proto: "tasks"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      switch fieldNumber {
      case 4032: try decoder.decodeSingularStringField(value: &self.name)
      case 4947: try decoder.decodeSingularBytesField(value: &self.id)
      case 8856: try decoder.decodeRepeatedMessageField(value: &self.tasks)
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.name.isEmpty {
      try visitor.visitSingularStringField(value: self.name, fieldNumber: 4032)
    }
    if !self.id.isEmpty {
      try visitor.visitSingularBytesField(value: self.id, fieldNumber: 4947)
    }
    if !self.tasks.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.tasks, fieldNumber: 8856)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: ViewModel.Home.Project, rhs: ViewModel.Home.Project) -> Bool {
    if lhs.id != rhs.id {return false}
    if lhs.name != rhs.name {return false}
    if lhs.tasks != rhs.tasks {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension ViewModel.Home.Project.Task: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = ViewModel.Home.Project.protoMessageName + ".Task"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    2946: .same(proto: "ID"),
    4336: .same(proto: "name"),
    2990: .same(proto: "emoji"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      switch fieldNumber {
      case 2946: try decoder.decodeSingularBytesField(value: &self.id)
      case 2990: try decoder.decodeSingularStringField(value: &self.emoji)
      case 4336: try decoder.decodeSingularStringField(value: &self.name)
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.id.isEmpty {
      try visitor.visitSingularBytesField(value: self.id, fieldNumber: 2946)
    }
    if !self.emoji.isEmpty {
      try visitor.visitSingularStringField(value: self.emoji, fieldNumber: 2990)
    }
    if !self.name.isEmpty {
      try visitor.visitSingularStringField(value: self.name, fieldNumber: 4336)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: ViewModel.Home.Project.Task, rhs: ViewModel.Home.Project.Task) -> Bool {
    if lhs.id != rhs.id {return false}
    if lhs.name != rhs.name {return false}
    if lhs.emoji != rhs.emoji {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
