import { useNavigate } from "react-router-dom";
import { api } from "../api";
import { LabelButton } from "../components/LabelButton";

export const Game = () => {
  const navigate = useNavigate();

  const handleNewEvent = () => {
    navigate("capture");
  };

  return (
    <div>
      <h1>Find the item below</h1>

      <p>Spoon</p>

      <LabelButton onClick={handleNewEvent} text="Found item" />
    </div>
  );
};
