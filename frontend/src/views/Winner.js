import axios from "axios";
import { ParticipantListItem } from "../components/ParticipantListItem";

export const Winner = () => {
  return (
    <div>
      <h1>Winner!</h1>

      <ParticipantListItem src="/images/avatar/1.jpg" name="Remy Sharp" />
    </div>
  );
};
