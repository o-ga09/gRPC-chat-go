import { websocketAtom } from "../state/websocket";
import { messageListAtom } from "../state/message";
import { useRecoilCallback, useRecoilValue } from "recoil";
import { Message } from "../models/message";

export const useMessageList = (): Message[] => {
  const socket = useRecoilValue(websocketAtom);
  const messageList = useRecoilValue(messageListAtom);

  const updateMessageList = useRecoilCallback(
    ({ set }) =>
      (message: Message) => {
        set(messageListAtom, [...messageList, message]);
      }
  );
  socket.onmessage = (msg) => {
    console.log(msg.data);
    const content = JSON.parse(msg.data as string);
    const message: Message = { 
      msgBody: content.msgBody,
      sender: content.sender,
      sendAt: content.sendAt
    };
    updateMessageList(message);
  };

  return messageList;
};