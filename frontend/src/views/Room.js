import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { api } from "../api";
import { ParticipantListItem } from "../components/ParticipantListItem";
import { IconLabelButton } from "../components/IconLabelButton";
import { LabelButton } from "../components/LabelButton";
import FormGroup from "@mui/material/FormGroup";
import FormControlLabel from "@mui/material/FormControlLabel";
import TextField from "@mui/material/TextField";
import Checkbox from "@mui/material/Checkbox";

export const Room = () => {
  const navigate = useNavigate();
  const [participants, setParticipants] = useState([]);

  const fetchParticipants = async () => {
    const response = await api.get("room");
    setParticipants(response.data);
  };

  const handlePlay = () => {
    navigate("game");
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

      <FormGroup sx={{ display: "flex", flexDirection: "column" }}>
        <TextField
          label="Stake amount"
          variant="outlined"
          type="number"
          margin="normal"
        />
        <FormControlLabel control={<Checkbox />} label="Ready to play" />
        <p style={{ fontSize: 13 }}>
          Check the box once you have entered your stake and are ready to play.
          Once all players have confirmed they are ready to play. You can hit
          the play button.
        </p>
      </FormGroup>

      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          padding: 5,
          gap: 5,
        }}
      >
        <IconLabelButton iconName="link" text="Copy invitation link" />
        <LabelButton onClick={handlePlay} text="Play game" />
      </div>
    </div>
  );
};
