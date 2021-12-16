import { getIn, Formik, Form, Field } from "formik";
import * as yup from "yup";
import { TextField } from "@material-ui/core";

const schema = yup.object({
  stake: yup
    .number("Enter a stake amount")
    .required("Stake amount is required")
    .min(1, "Minimum stake amount is £1.00"),
});

export const TextFormField = ({ field, form, ...props }) => {
  const errorText =
    getIn(form.touched, field.name) && getIn(form.errors, field.name);

  return (
    <TextField
      fullWidth
      margin="normal"
      helperText={errorText}
      error={!!errorText}
      {...field}
      {...props}
    />
  );
};

//   <Formik
//     validationSchema={schema}
//     initialValues={{ stake: 0 }}
//     onSubmit={handleSubmit}
//   >
//     {() => (
//       <Form>
//         <Field
//           label="Stake amount"
//           name="stake"
//           component={TextFormField}
//         />
//         <Field
//           label="Ready to play"
//           name="ready"
//           component={CheckboxFormField}
//         />
//       </Form>
//     )}
//   </Formik>
