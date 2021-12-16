import Button from "@mui/material/Button";

export const LabelButton = ({ text, ...props }) => {
  return (
    <Button
      variant="contained"
      sx={{
        minWidth: 167,
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
