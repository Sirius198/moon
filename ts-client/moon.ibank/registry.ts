import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgReceive } from "./types/moon/ibank/tx";
import { MsgSend } from "./types/moon/ibank/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/moon.ibank.MsgReceive", MsgReceive],
    ["/moon.ibank.MsgSend", MsgSend],
    
];

export { msgTypes }