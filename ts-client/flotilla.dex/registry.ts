import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSendBuyOrder } from "./types/dex/tx";
import { MsgSendCreatePair } from "./types/dex/tx";
import { MsgCancelBuyOrder } from "./types/dex/tx";
import { MsgSendSellOrder } from "./types/dex/tx";
import { MsgCancelSellOrder } from "./types/dex/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/flotilla.dex.MsgSendBuyOrder", MsgSendBuyOrder],
    ["/flotilla.dex.MsgSendCreatePair", MsgSendCreatePair],
    ["/flotilla.dex.MsgCancelBuyOrder", MsgCancelBuyOrder],
    ["/flotilla.dex.MsgSendSellOrder", MsgSendSellOrder],
    ["/flotilla.dex.MsgCancelSellOrder", MsgCancelSellOrder],
    
];

export { msgTypes }