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
      sendAt: ToDate()
    };
    socket.send(JSON.stringify(message));
    setInput("");
  }, [input, socket]);

  return { input, setInput, send };
};

function ToDate() :string {
  const date = new Date();

  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, "0");
  const day = String(date.getDate()).padStart(2, "0");
  const hours = String(date.getHours()).padStart(2, "0");
  const minutes = String(date.getMinutes()).padStart(2, "0");
  const seconds = String(date.getSeconds()).padStart(2, "0");

  const customFormat = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;

  return customFormat;
}