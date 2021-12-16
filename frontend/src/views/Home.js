import { Link } from "react-router-dom";
import { LabelButton } from "../components/LabelButton";

export const Home = () => {
  return (
    <div>
      <h1>Meta betting</h1>

      <Link to="/room">
        <LabelButton text="Create new event" />
      </Link>
    </div>
  );
};
