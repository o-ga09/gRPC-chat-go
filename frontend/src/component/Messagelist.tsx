import { useMessageList } from "../hooks/useMessagelist";

export const MessageList = () => {
  const messageList = useMessageList();

  return (
    <div>
        <table border={1}>
          <tr>
            <th>user</th>
            <th>msg</th>
            <th>date</th>
          </tr>
          {messageList.map((m, i) => (
          <tr>
            <td key={i}>{m.sender}</td>
            <td key={i}>{m.msgBody}</td>
            <td key={i}>{m.sendAt}</td>
          </tr>
          ))}
        </table>
    </div>
  );
};