import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { api } from "../api";
import { ParticipantListItem } from "../components/ParticipantListItem";

export const Room = () => {
  const { id } = useParams();
  const [participants, setParticipants] = useState([]);

  const fetchParticipants = async () => {
    const response = await api.get("room");
    console.log(response);
    setParticipants(response.data);
  };

  useEffect(() => {
    fetchParticipants();
  }, []);

  return (
    <div>
      <h2>Participants</h2>

      {participants.map((player) => (
        <ParticipantListItem src={player.img} name={player.name} />
      ))}

      <button>Copy invitation link</button>
    </div>
  );
};
