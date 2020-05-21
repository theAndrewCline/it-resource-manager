CREATE TABLE owners ( 
    id UUID PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE computers ( 
    id UUID PRIMARY KEY,
    owner_id UUID,
    description TEXT NOT NULL    
);


CREATE TABLE parts ( 
    id UUID PRIMARY KEY,
    computer_id UUID,
    name TEXT NOT NULL,
    model_number TEXT NOT NULL
);

