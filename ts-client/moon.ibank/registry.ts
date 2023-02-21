import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSend } from "./types/moon/ibank/tx";
import { MsgReceive } from "./types/moon/ibank/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/moon.ibank.MsgSend", MsgSend],
    ["/moon.ibank.MsgReceive", MsgReceive],
    
];

export { msgTypes }