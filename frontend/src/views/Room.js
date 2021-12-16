import * as React from 'react';
import {ParticipantListItem} from "../components/ParticipantListItem";


export const Room = () => {
  return (
    <div>
      <h2>Participants</h2>

        <ParticipantListItem src="/images/avatar/1.jpg" name="Remy Sharp"/>
        <ParticipantListItem src="/images/avatar/2.jpg" name="Travis Howard"/>
        <ParticipantListItem src="/images/avatar/3.jpg" name="Cindy Baker"/>
        <ParticipantListItem src="/images/avatar/4.jpg" name="Agnes Walker"/>
        <ParticipantListItem src="/images/avatar/5.jpg" name="Trevor Henderson"/>

      <span>
        ~Loop over mocked participants here and render components for each~
      </span>

      <button>Copy invitation link</button>
    </div>
  );
};
