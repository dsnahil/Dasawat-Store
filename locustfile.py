import random
from locust import HttpUser, task, between

class ProductUser(HttpUser):
    # Wait 1 to 5 seconds between tasks
    wait_time = between(1, 5)

    @task(10)  # Make GET requests 10 times more likely than POST requests
    def get_product(self):
        # Pick a random product ID to request
        product_id = random.randint(1, 500)
        self.client.get(f"/products/{product_id}")

    @task(1)   # Make POST requests less frequent
    def create_product(self):
        product_id = random.randint(1, 500)
        # Define the JSON data to send
        product_data = {
          "product_id": product_id,
          "sku": f"LOCUST-TEST-{product_id}",
          "manufacturer": "Load Test Inc.",
          "category_id": 1,
          "weight": 200,
          "some_other_id": 300
        }
        self.client.post(f"/products/{product_id}/details", json=product_data)