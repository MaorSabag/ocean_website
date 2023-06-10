import React from 'react';
import { Alert } from '@mui/material';
import { errorMessage } from '../../Models/index';

type Props = {
  errorMessage: errorMessage | undefined;
  isOpenAlert: boolean;
  onClose: () => void;
};

export const AlertPopUp = (props: Props) => {
  const { errorMessage, isOpenAlert, onClose } = props;
  console.log("Got in alertPopUp");
  return (
    <div className="centered-container">
      <Alert
        className="centered-alert"
        severity="error"
        variant="filled"
        onClose={onClose}
        title={errorMessage?.Status.toString()}
      >
        {errorMessage?.Error.toString()}
      </Alert>
    </div>
  );
};
