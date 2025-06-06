// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: outputs/speed.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace ProtobufMsgs {

  /// <summary>Holder for reflection information generated from outputs/speed.proto</summary>
  public static partial class SpeedReflection {

    #region Descriptor
    /// <summary>File descriptor for outputs/speed.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static SpeedReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChNvdXRwdXRzL3NwZWVkLnByb3RvEg1wcm90b2J1Zl9tc2dzIiAKEVNwZWVk",
            "U2Vuc29yT3V0cHV0EgsKA3JwbRgBIAEoBUIQWg5hc2UvcGJfb3V0cHV0c2IG",
            "cHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::ProtobufMsgs.SpeedSensorOutput), global::ProtobufMsgs.SpeedSensorOutput.Parser, new[]{ "Rpm" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class SpeedSensorOutput : pb::IMessage<SpeedSensorOutput> {
    private static readonly pb::MessageParser<SpeedSensorOutput> _parser = new pb::MessageParser<SpeedSensorOutput>(() => new SpeedSensorOutput());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<SpeedSensorOutput> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::ProtobufMsgs.SpeedReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SpeedSensorOutput() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SpeedSensorOutput(SpeedSensorOutput other) : this() {
      rpm_ = other.rpm_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SpeedSensorOutput Clone() {
      return new SpeedSensorOutput(this);
    }

    /// <summary>Field number for the "rpm" field.</summary>
    public const int RpmFieldNumber = 1;
    private int rpm_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int Rpm {
      get { return rpm_; }
      set {
        rpm_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as SpeedSensorOutput);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(SpeedSensorOutput other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Rpm != other.Rpm) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Rpm != 0) hash ^= Rpm.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Rpm != 0) {
        output.WriteRawTag(8);
        output.WriteInt32(Rpm);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Rpm != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Rpm);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(SpeedSensorOutput other) {
      if (other == null) {
        return;
      }
      if (other.Rpm != 0) {
        Rpm = other.Rpm;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 8: {
            Rpm = input.ReadInt32();
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code
