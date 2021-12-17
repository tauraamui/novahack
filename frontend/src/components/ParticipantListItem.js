import * as React from "react";
import Avatar from "@mui/material/Avatar";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemText from "@mui/material/ListItemText";

export const ParticipantListItem = ({ src, name, ...props }) => {
  return (
    <ListItemButton {...props}>
      <Avatar sx={{ marginRight: 2 }} alt={name} src={src} />
      <ListItemText primary={name} />
    </ListItemButton>
  );
};
