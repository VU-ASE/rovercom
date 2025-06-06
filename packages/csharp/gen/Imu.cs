// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: outputs/imu.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace ProtobufMsgs {

  /// <summary>Holder for reflection information generated from outputs/imu.proto</summary>
  public static partial class ImuReflection {

    #region Descriptor
    /// <summary>File descriptor for outputs/imu.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static ImuReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChFvdXRwdXRzL2ltdS5wcm90bxINcHJvdG9idWZfbXNncyLIAwoPSW11U2Vu",
            "c29yT3V0cHV0EhMKC3RlbXBlcmF0dXJlGAEgASgFEjsKDG1hZ25ldG9tZXRl",
            "chgCIAEoCzIlLnByb3RvYnVmX21zZ3MuSW11U2Vuc29yT3V0cHV0LlZlY3Rv",
            "chI4CglneXJvc2NvcGUYAyABKAsyJS5wcm90b2J1Zl9tc2dzLkltdVNlbnNv",
            "ck91dHB1dC5WZWN0b3ISNAoFZXVsZXIYBCABKAsyJS5wcm90b2J1Zl9tc2dz",
            "LkltdVNlbnNvck91dHB1dC5WZWN0b3ISPAoNYWNjZWxlcm9tZXRlchgFIAEo",
            "CzIlLnByb3RvYnVmX21zZ3MuSW11U2Vuc29yT3V0cHV0LlZlY3RvchJCChNs",
            "aW5lYXJBY2NlbGVyb21ldGVyGAYgASgLMiUucHJvdG9idWZfbXNncy5JbXVT",
            "ZW5zb3JPdXRwdXQuVmVjdG9yEjcKCHZlbG9jaXR5GAcgASgLMiUucHJvdG9i",
            "dWZfbXNncy5JbXVTZW5zb3JPdXRwdXQuVmVjdG9yEg0KBXNwZWVkGAggASgC",
            "GikKBlZlY3RvchIJCgF4GAEgASgCEgkKAXkYAiABKAISCQoBehgDIAEoAkIQ",
            "Wg5hc2UvcGJfb3V0cHV0c2IGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::ProtobufMsgs.ImuSensorOutput), global::ProtobufMsgs.ImuSensorOutput.Parser, new[]{ "Temperature", "Magnetometer", "Gyroscope", "Euler", "Accelerometer", "LinearAccelerometer", "Velocity", "Speed" }, null, null, null, new pbr::GeneratedClrTypeInfo[] { new pbr::GeneratedClrTypeInfo(typeof(global::ProtobufMsgs.ImuSensorOutput.Types.Vector), global::ProtobufMsgs.ImuSensorOutput.Types.Vector.Parser, new[]{ "X", "Y", "Z" }, null, null, null, null)})
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class ImuSensorOutput : pb::IMessage<ImuSensorOutput> {
    private static readonly pb::MessageParser<ImuSensorOutput> _parser = new pb::MessageParser<ImuSensorOutput>(() => new ImuSensorOutput());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<ImuSensorOutput> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::ProtobufMsgs.ImuReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ImuSensorOutput() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ImuSensorOutput(ImuSensorOutput other) : this() {
      temperature_ = other.temperature_;
      magnetometer_ = other.magnetometer_ != null ? other.magnetometer_.Clone() : null;
      gyroscope_ = other.gyroscope_ != null ? other.gyroscope_.Clone() : null;
      euler_ = other.euler_ != null ? other.euler_.Clone() : null;
      accelerometer_ = other.accelerometer_ != null ? other.accelerometer_.Clone() : null;
      linearAccelerometer_ = other.linearAccelerometer_ != null ? other.linearAccelerometer_.Clone() : null;
      velocity_ = other.velocity_ != null ? other.velocity_.Clone() : null;
      speed_ = other.speed_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ImuSensorOutput Clone() {
      return new ImuSensorOutput(this);
    }

    /// <summary>Field number for the "temperature" field.</summary>
    public const int TemperatureFieldNumber = 1;
    private int temperature_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int Temperature {
      get { return temperature_; }
      set {
        temperature_ = value;
      }
    }

    /// <summary>Field number for the "magnetometer" field.</summary>
    public const int MagnetometerFieldNumber = 2;
    private global::ProtobufMsgs.ImuSensorOutput.Types.Vector magnetometer_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::ProtobufMsgs.ImuSensorOutput.Types.Vector Magnetometer {
      get { return magnetometer_; }
      set {
        magnetometer_ = value;
      }
    }

    /// <summary>Field number for the "gyroscope" field.</summary>
    public const int GyroscopeFieldNumber = 3;
    private global::ProtobufMsgs.ImuSensorOutput.Types.Vector gyroscope_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::ProtobufMsgs.ImuSensorOutput.Types.Vector Gyroscope {
      get { return gyroscope_; }
      set {
        gyroscope_ = value;
      }
    }

    /// <summary>Field number for the "euler" field.</summary>
    public const int EulerFieldNumber = 4;
    private global::ProtobufMsgs.ImuSensorOutput.Types.Vector euler_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::ProtobufMsgs.ImuSensorOutput.Types.Vector Euler {
      get { return euler_; }
      set {
        euler_ = value;
      }
    }

    /// <summary>Field number for the "accelerometer" field.</summary>
    public const int AccelerometerFieldNumber = 5;
    private global::ProtobufMsgs.ImuSensorOutput.Types.Vector accelerometer_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::ProtobufMsgs.ImuSensorOutput.Types.Vector Accelerometer {
      get { return accelerometer_; }
      set {
        accelerometer_ = value;
      }
    }

    /// <summary>Field number for the "linearAccelerometer" field.</summary>
    public const int LinearAccelerometerFieldNumber = 6;
    private global::ProtobufMsgs.ImuSensorOutput.Types.Vector linearAccelerometer_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::ProtobufMsgs.ImuSensorOutput.Types.Vector LinearAccelerometer {
      get { return linearAccelerometer_; }
      set {
        linearAccelerometer_ = value;
      }
    }

    /// <summary>Field number for the "velocity" field.</summary>
    public const int VelocityFieldNumber = 7;
    private global::ProtobufMsgs.ImuSensorOutput.Types.Vector velocity_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::ProtobufMsgs.ImuSensorOutput.Types.Vector Velocity {
      get { return velocity_; }
      set {
        velocity_ = value;
      }
    }

    /// <summary>Field number for the "speed" field.</summary>
    public const int SpeedFieldNumber = 8;
    private float speed_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public float Speed {
      get { return speed_; }
      set {
        speed_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as ImuSensorOutput);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(ImuSensorOutput other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Temperature != other.Temperature) return false;
      if (!object.Equals(Magnetometer, other.Magnetometer)) return false;
      if (!object.Equals(Gyroscope, other.Gyroscope)) return false;
      if (!object.Equals(Euler, other.Euler)) return false;
      if (!object.Equals(Accelerometer, other.Accelerometer)) return false;
      if (!object.Equals(LinearAccelerometer, other.LinearAccelerometer)) return false;
      if (!object.Equals(Velocity, other.Velocity)) return false;
      if (!pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.Equals(Speed, other.Speed)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Temperature != 0) hash ^= Temperature.GetHashCode();
      if (magnetometer_ != null) hash ^= Magnetometer.GetHashCode();
      if (gyroscope_ != null) hash ^= Gyroscope.GetHashCode();
      if (euler_ != null) hash ^= Euler.GetHashCode();
      if (accelerometer_ != null) hash ^= Accelerometer.GetHashCode();
      if (linearAccelerometer_ != null) hash ^= LinearAccelerometer.GetHashCode();
      if (velocity_ != null) hash ^= Velocity.GetHashCode();
      if (Speed != 0F) hash ^= pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.GetHashCode(Speed);
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
      if (Temperature != 0) {
        output.WriteRawTag(8);
        output.WriteInt32(Temperature);
      }
      if (magnetometer_ != null) {
        output.WriteRawTag(18);
        output.WriteMessage(Magnetometer);
      }
      if (gyroscope_ != null) {
        output.WriteRawTag(26);
        output.WriteMessage(Gyroscope);
      }
      if (euler_ != null) {
        output.WriteRawTag(34);
        output.WriteMessage(Euler);
      }
      if (accelerometer_ != null) {
        output.WriteRawTag(42);
        output.WriteMessage(Accelerometer);
      }
      if (linearAccelerometer_ != null) {
        output.WriteRawTag(50);
        output.WriteMessage(LinearAccelerometer);
      }
      if (velocity_ != null) {
        output.WriteRawTag(58);
        output.WriteMessage(Velocity);
      }
      if (Speed != 0F) {
        output.WriteRawTag(69);
        output.WriteFloat(Speed);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Temperature != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Temperature);
      }
      if (magnetometer_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Magnetometer);
      }
      if (gyroscope_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Gyroscope);
      }
      if (euler_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Euler);
      }
      if (accelerometer_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Accelerometer);
      }
      if (linearAccelerometer_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(LinearAccelerometer);
      }
      if (velocity_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Velocity);
      }
      if (Speed != 0F) {
        size += 1 + 4;
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(ImuSensorOutput other) {
      if (other == null) {
        return;
      }
      if (other.Temperature != 0) {
        Temperature = other.Temperature;
      }
      if (other.magnetometer_ != null) {
        if (magnetometer_ == null) {
          Magnetometer = new global::ProtobufMsgs.ImuSensorOutput.Types.Vector();
        }
        Magnetometer.MergeFrom(other.Magnetometer);
      }
      if (other.gyroscope_ != null) {
        if (gyroscope_ == null) {
          Gyroscope = new global::ProtobufMsgs.ImuSensorOutput.Types.Vector();
        }
        Gyroscope.MergeFrom(other.Gyroscope);
      }
      if (other.euler_ != null) {
        if (euler_ == null) {
          Euler = new global::ProtobufMsgs.ImuSensorOutput.Types.Vector();
        }
        Euler.MergeFrom(other.Euler);
      }
      if (other.accelerometer_ != null) {
        if (accelerometer_ == null) {
          Accelerometer = new global::ProtobufMsgs.ImuSensorOutput.Types.Vector();
        }
        Accelerometer.MergeFrom(other.Accelerometer);
      }
      if (other.linearAccelerometer_ != null) {
        if (linearAccelerometer_ == null) {
          LinearAccelerometer = new global::ProtobufMsgs.ImuSensorOutput.Types.Vector();
        }
        LinearAccelerometer.MergeFrom(other.LinearAccelerometer);
      }
      if (other.velocity_ != null) {
        if (velocity_ == null) {
          Velocity = new global::ProtobufMsgs.ImuSensorOutput.Types.Vector();
        }
        Velocity.MergeFrom(other.Velocity);
      }
      if (other.Speed != 0F) {
        Speed = other.Speed;
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
            Temperature = input.ReadInt32();
            break;
          }
          case 18: {
            if (magnetometer_ == null) {
              Magnetometer = new global::ProtobufMsgs.ImuSensorOutput.Types.Vector();
            }
            input.ReadMessage(Magnetometer);
            break;
          }
          case 26: {
            if (gyroscope_ == null) {
              Gyroscope = new global::ProtobufMsgs.ImuSensorOutput.Types.Vector();
            }
            input.ReadMessage(Gyroscope);
            break;
          }
          case 34: {
            if (euler_ == null) {
              Euler = new global::ProtobufMsgs.ImuSensorOutput.Types.Vector();
            }
            input.ReadMessage(Euler);
            break;
          }
          case 42: {
            if (accelerometer_ == null) {
              Accelerometer = new global::ProtobufMsgs.ImuSensorOutput.Types.Vector();
            }
            input.ReadMessage(Accelerometer);
            break;
          }
          case 50: {
            if (linearAccelerometer_ == null) {
              LinearAccelerometer = new global::ProtobufMsgs.ImuSensorOutput.Types.Vector();
            }
            input.ReadMessage(LinearAccelerometer);
            break;
          }
          case 58: {
            if (velocity_ == null) {
              Velocity = new global::ProtobufMsgs.ImuSensorOutput.Types.Vector();
            }
            input.ReadMessage(Velocity);
            break;
          }
          case 69: {
            Speed = input.ReadFloat();
            break;
          }
        }
      }
    }

    #region Nested types
    /// <summary>Container for nested types declared in the ImuSensorOutput message type.</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static partial class Types {
      public sealed partial class Vector : pb::IMessage<Vector> {
        private static readonly pb::MessageParser<Vector> _parser = new pb::MessageParser<Vector>(() => new Vector());
        private pb::UnknownFieldSet _unknownFields;
        [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
        public static pb::MessageParser<Vector> Parser { get { return _parser; } }

        [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
        public static pbr::MessageDescriptor Descriptor {
          get { return global::ProtobufMsgs.ImuSensorOutput.Descriptor.NestedTypes[0]; }
        }

        [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
        pbr::MessageDescriptor pb::IMessage.Descriptor {
          get { return Descriptor; }
        }

        [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
        public Vector() {
          OnConstruction();
        }

        partial void OnConstruction();

        [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
        public Vector(Vector other) : this() {
          x_ = other.x_;
          y_ = other.y_;
          z_ = other.z_;
          _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
        }

        [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
        public Vector Clone() {
          return new Vector(this);
        }

        /// <summary>Field number for the "x" field.</summary>
        public const int XFieldNumber = 1;
        private float x_;
        [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
        public float X {
          get { return x_; }
          set {
            x_ = value;
          }
        }

        /// <summary>Field number for the "y" field.</summary>
        public const int YFieldNumber = 2;
        private float y_;
        [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
        public float Y {
          get { return y_; }
          set {
            y_ = value;
          }
        }

        /// <summary>Field number for the "z" field.</summary>
        public const int ZFieldNumber = 3;
        private float z_;
        [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
        public float Z {
          get { return z_; }
          set {
            z_ = value;
          }
        }

        [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
        public override bool Equals(object other) {
          return Equals(other as Vector);
        }

        [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
        public bool Equals(Vector other) {
          if (ReferenceEquals(other, null)) {
            return false;
          }
          if (ReferenceEquals(other, this)) {
            return true;
          }
          if (!pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.Equals(X, other.X)) return false;
          if (!pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.Equals(Y, other.Y)) return false;
          if (!pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.Equals(Z, other.Z)) return false;
          return Equals(_unknownFields, other._unknownFields);
        }

        [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
        public override int GetHashCode() {
          int hash = 1;
          if (X != 0F) hash ^= pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.GetHashCode(X);
          if (Y != 0F) hash ^= pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.GetHashCode(Y);
          if (Z != 0F) hash ^= pbc::ProtobufEqualityComparers.BitwiseSingleEqualityComparer.GetHashCode(Z);
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
          if (X != 0F) {
            output.WriteRawTag(13);
            output.WriteFloat(X);
          }
          if (Y != 0F) {
            output.WriteRawTag(21);
            output.WriteFloat(Y);
          }
          if (Z != 0F) {
            output.WriteRawTag(29);
            output.WriteFloat(Z);
          }
          if (_unknownFields != null) {
            _unknownFields.WriteTo(output);
          }
        }

        [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
        public int CalculateSize() {
          int size = 0;
          if (X != 0F) {
            size += 1 + 4;
          }
          if (Y != 0F) {
            size += 1 + 4;
          }
          if (Z != 0F) {
            size += 1 + 4;
          }
          if (_unknownFields != null) {
            size += _unknownFields.CalculateSize();
          }
          return size;
        }

        [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
        public void MergeFrom(Vector other) {
          if (other == null) {
            return;
          }
          if (other.X != 0F) {
            X = other.X;
          }
          if (other.Y != 0F) {
            Y = other.Y;
          }
          if (other.Z != 0F) {
            Z = other.Z;
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
              case 13: {
                X = input.ReadFloat();
                break;
              }
              case 21: {
                Y = input.ReadFloat();
                break;
              }
              case 29: {
                Z = input.ReadFloat();
                break;
              }
            }
          }
        }

      }

    }
    #endregion

  }

  #endregion

}

#endregion Designer generated code
