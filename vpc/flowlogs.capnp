using Go = import "/go.capnp";
@0x9f5a5320cda72e66;

$Go.package("flowlogs");
$Go.import("vpc/flowlogs");

struct VpcFlowLogEntry {
  version @0 :Int8;
  accountId @1 :UInt64;
  interfaceId @2 :Text;
  srcAddr @3 :UInt32;
  dstAddr @4 :UInt32;
  srcPort @5 :UInt16;
  dstPort @6 :UInt16;
  protocol @7 :UInt8;
  packets @8 :UInt16;
  bytes @9 :UInt64;
  start @10 :UInt32;
  end @11 :UInt32;
  action @12 :Action;
  logStatus @13 :LogStatus;

  enum Action {
    accept @0;
    reject @1;
  }

  enum LogStatus {
    ok @0;
    noData @1;
    skipData @2;
  }
}
