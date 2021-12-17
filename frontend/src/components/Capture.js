import React from "react";
import Webcam from "react-webcam";
import { LabelButton } from "./LabelButton";

const videoConstraints = {
  width: 640,
  height: 480,
  facingMode: "user",
};

export const Capture = () => {
  const webcamRef = React.useRef(null);

  const capture = React.useCallback(() => {
    const imageSrc = webcamRef.current.getScreenshot();
    console.log(imageSrc);
  }, [webcamRef]);

  return (
    <>
      <Webcam
        audio={false}
        height={480}
        ref={webcamRef}
        screenshotFormat="image/jpeg"
        width={640}
        videoConstraints={videoConstraints}
      />
      <LabelButton onClick={capture} text="Capture photo" />
    </>
  );
};
