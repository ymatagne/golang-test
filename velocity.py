
import json

with open("./input.txt", "r") as ins:
    clients={}
    for line in ins:
        data = json.loads(line)
        if data['customer_id'] not in clients:
            clients[data['customer_id']] = {'performed_per_day' : 0, 'total_per_day': 0,'total_per_week': 0}

        customer_state=clients[data['customer_id']]['performed_per_day']
        customer_state['performed_per_day']=customer_state['performed_per_day']+1
        

        clients[data['customer_id']]['performed_per_day']=
    print(clients)
