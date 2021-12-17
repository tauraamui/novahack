import { useState, useEffect } from "react";
import "./StopWatch.css";
import { Timer } from "./Timer";

export const StopWatch = () => {
  const [isActive, setIsActive] = useState(false);
  const [time, setTime] = useState(0);

  useEffect(() => {
    let interval = null;

    if (isActive === false) {
      interval = setInterval(() => {
        setTime((time) => time + 10);
      }, 10);
    } else {
      clearInterval(interval);
    }
    return () => {
      clearInterval(interval);
    };
  }, [isActive]);

  return (
    <div className="stop-watch">
      <Timer time={time} />
    </div>
  );
};
