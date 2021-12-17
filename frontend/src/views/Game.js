import { useNavigate } from "react-router-dom";
import { LabelButton } from "../components/LabelButton";
import { StopWatch } from "../components/StopWatch";

export const Game = () => {
  const navigate = useNavigate();

  const handleNewEvent = () => {
    navigate("capture");
  };

  return (
    <div>
      <h1>Find the item below</h1>

      <p>Spoon</p>

      <StopWatch />

      <LabelButton onClick={handleNewEvent} text="Found item" />
    </div>
  );
};
