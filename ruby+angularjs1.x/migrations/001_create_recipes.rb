Sequel.migration do
  change do
    create_table :recipes do
      primary_key :id
      String :title, unique: true
      String :description, text: true
      String :instructions, text: true
      jsonb :ingredients
    end
  end
end
