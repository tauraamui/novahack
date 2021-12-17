from flask import Flask, request, jsonify
from PIL import Image
import json
import tensorflow as tf
from tensorflow import keras
from tensorflow.keras import layers
from tensorflow.keras.models import Sequential
import base64
import logging
import numpy as np

app = Flask(__name__)


@app.route("/get_predictions", methods=['GET'])
def predict():
    def load_model():
        data_augmentation = keras.Sequential([
            layers.RandomFlip("horizontal", input_shape=(180, 180, 3)),
            layers.RandomRotation(0.1),
            layers.RandomZoom(0.1),
        ])

        num_classes = 4
        my_model = Sequential([
            data_augmentation,
            layers.Rescaling(1. / 255),
            layers.Conv2D(16, 3, padding='same', activation='relu'),
            layers.MaxPooling2D(),
            layers.Conv2D(32, 3, padding='same', activation='relu'),
            layers.MaxPooling2D(),
            layers.Conv2D(64, 3, padding='same', activation='relu'),
            layers.MaxPooling2D(),
            layers.Dropout(0.3),
            layers.Flatten(),
            layers.Dense(128, activation='relu'),
            layers.Dense(num_classes)
        ])

        my_model.compile(optimizer='adam',
                         loss=tf.keras.losses.SparseCategoricalCrossentropy(
                             from_logits=True),
                         metrics=['accuracy'])

        my_model.load_weights('CV_API/weights.h5')
        return my_model

    try:
        class_names = ['backpack', 'binoculars', 'calculator', 'flashlight']
        model = load_model()

        img = tf.keras.utils.load_img('CV_API/data/backpack_test.jpg',
                                      target_size=(180, 180))

        img_array = tf.keras.utils.img_to_array(img)
        img_array = tf.expand_dims(img_array, 0)  # Create a batch
        predictions = model.predict(img_array)
        score = tf.nn.softmax(predictions[0])

        response = {
            'class': str(class_names[np.argmax(score)]),
            'confidence': round(float(np.max(score)), 2)
        }

        return response, 200

    except Exception as e:
        logging.error(e)
        return "Internal Server Error", 500
