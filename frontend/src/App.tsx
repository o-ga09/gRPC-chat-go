import { MessageInput } from "./component/MessageInput";
import { MessageList } from "./component/Messagelist";

export const App = () => {
  return (
    <div>
      <h1>Simple Chat</h1>
      <MessageInput />
      <MessageList />
    </div>
  );
};