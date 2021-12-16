import { Routes, Route } from "react-router-dom";
import { Home } from "./views/Home";
import { Room } from "./views/Room";
import { CaptureImage } from "./views/CaptureImage";
import { Stake } from "./views/Stake";

export const App = () => {
  return (
    <div>
      <Routes>
        <Route path="" element={<Home />} />
        <Route path="room/:id" element={<Room />} />
        <Route path="capture" element={<CaptureImage />} />
        <Route path="stake" element={<Stake />} />
      </Routes>
    </div>
  );
};
