import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSendSellOrder } from "./types/dex/tx";
import { MsgSendBuyOrder } from "./types/dex/tx";
import { MsgSendCreatePair } from "./types/dex/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/flotilla.dex.MsgSendSellOrder", MsgSendSellOrder],
    ["/flotilla.dex.MsgSendBuyOrder", MsgSendBuyOrder],
    ["/flotilla.dex.MsgSendCreatePair", MsgSendCreatePair],
    
];

export { msgTypes }