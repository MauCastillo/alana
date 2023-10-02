import numpy as np
from tensorflow import keras

# Generate random training data
num_samples = 1000
input_data = np.random.rand(num_samples, 2)
target_data = np.sum(input_data, axis=1)
print("input_data", input_data)

layer = keras.layers.Dense(16, activation='relu', input_shape=(2,)),
# Create a sequential model
model = keras.Sequential([
    keras.layers.Dense(16, activation='relu', input_shape=(2,)),
    keras.layers.Dense(1)
])

# Compile the model
model.compile(optimizer='adam', loss='mse')

# Train the model
model.fit(input_data, target_data, epochs=10, batch_size=32)

# Test the model
test_data = np.array([[0.1, 0.2], [0.3, 0.4]])
predictions = model.predict(test_data)
print(layer.get_weights())
print(predictions)