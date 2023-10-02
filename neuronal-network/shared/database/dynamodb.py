import boto3
import numpy as np


class Dynamodb:
    FLOAT_PRECISION = 2
    input_data = []
    target_data = []

    def __init__(self, table, pair):
        self.Table = table
        self.Pair = pair

    def read_table_filter(self):
        dynamodb = boto3.client('dynamodb')
        filter_expression = '#name = :value'
        expression_attribute_names = {'#name': 'name'}

        expression_attribute_values = {':value': {'S': self.Pair}}

        response = dynamodb.scan(
            TableName=self.Table,
            FilterExpression=filter_expression,
            ExpressionAttributeNames=expression_attribute_names,
            ExpressionAttributeValues=expression_attribute_values
        )

        return response['Items']

    def read_table_raw(self):
        dynamodb = boto3.client('dynamodb')

        response = dynamodb.scan(TableName=self.Table)

        return response['Items']
