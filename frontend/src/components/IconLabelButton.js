import Button from "@mui/material/Button";
import EditIcon from "@mui/icons-material/Edit";
import AddCircleIcon from "@mui/icons-material/AddCircle";

const buttonIcons = {
  add: <AddCircleIcon />,
  edit: <EditIcon />,
};

export const IconLabelButton = ({ iconName, text, ...props }) => {
  return (
    <Button
      variant="contained"
      startIcon={buttonIcons[iconName]}
      sx={{
        borderRadius: 32,
        color: "black",
        background: "#61A0FF",
        textTransform: "capitalize",
      }}
      {...props}
    >
      {text}
    </Button>
  );
};
