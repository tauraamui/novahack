import * as React from "react";
import FormControl from "@mui/material/FormControl";
import FormLabel from "@mui/material/FormLabel";
import RadioGroup from "@mui/material/RadioGroup";
import Radio from "@mui/material/Radio";
import Button from "@mui/material/Button";
import FormControlLabel from "@mui/material/FormControlLabel";

export const SelectStake = () => {
  return (
    <FormControl component="fieldset">
      <FormLabel component="legend">Select Stake</FormLabel>
      <RadioGroup row aria-label="gender" name="row-radio-buttons-group">
        <FormControlLabel value="5" control={<Radio />} label="5£" />
        <FormControlLabel value="10" control={<Radio />} label="10£" />
        <FormControlLabel value="15" control={<Radio />} label="15£" />
        <FormControlLabel value="20" control={<Radio />} label="20£" />
      </RadioGroup>
      <Button variant="contained">Submit</Button>
    </FormControl>
  );
};
