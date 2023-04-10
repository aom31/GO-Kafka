Producer side [

1.API to controller -Producer

    -have request topic
    
    +OpenAccountCommand
    
    +DepositFundCommand
    
    +WithdrawFundCommand
    
    +CloseAccountCommand
    
2. Service process 

3. producer to kafka via event

    Topic Event Kafka
    
    -OpenAccountEvent
    
    -DepositFundEvent
    
    -WithdrawFundEvent
    
    -CloseAccountEvent

]
Consumer side [

    1.Consumer Group subscribe Event kafka
    
    2. Handler handle data send to 
    
    3. Repository to database
    
]
