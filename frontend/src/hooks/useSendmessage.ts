import { useCallback, useState } from "react";
import { websocketAtom } from "../state/websocket";
import { useRecoilValue } from "recoil";
import { Message } from "../models/message";

export const useSendMessage = () => {
  const socket = useRecoilValue(websocketAtom);
  const [input, setInput] = useState<string>("");

  const send = useCallback(() => {
    if (input.length === 0) return;
    const message: Message = { 
      msgBody: input,
      sender: "user1",
      sendAt: "2023-09-11 23:59:59"
    };
    socket.send(JSON.stringify(message));
    setInput("");
  }, [input, socket]);

  return { input, setInput, send };
};