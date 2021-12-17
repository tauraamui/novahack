import { useNavigate } from "react-router-dom";
import { api } from "../api";
import { LabelButton } from "../components/LabelButton";

export const Home = () => {
  const navigate = useNavigate();

  const handleNewEvent = async () => {
    const response = await api.post("room");
    const roomId = response.data.room;
    navigate(`/room/${roomId}`);
  };

  return (
    <div>
      <h1>Meta betting</h1>

      <LabelButton onClick={handleNewEvent} text="Create new event" />
    </div>
  );
};
