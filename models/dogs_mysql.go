package models

import (
	"context"
	"time"
)

func (d *mysqlRepository) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `select id, breed, weight_low_lbs, weight_high_lbs, 
				cast(((weight_high_lbs+weight_low_lbs) / 2 ) as unsigned) as avg_weight,
				lifespan, coalesce(details,''),
				coalesce(alternate_names,''),coalesce(geographic_origin,'')
				from dog_breeds order by breed`

	var breeds []*DogBreed = make([]*DogBreed, 0)

	rows, err := d.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var b DogBreed
		err := rows.Scan(&b.Id, &b.Breed, &b.WeightLowLbs, &b.WeightHighLbs,
			&b.AverageWeight, &b.Lifespan, &b.Details, &b.AlternateNames,
			&b.GeographicOrigin)
		if err != nil {
			return nil, err
		}
		breeds = append(breeds, &b)
	}
	return breeds, nil
}

func (d *mysqlRepository) GetBreedByName(b string) (*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `select id, breed, weight_low_lbs, weight_high_lbs, 
				cast(((weight_high_lbs+weight_low_lbs) / 2 ) as unsigned) as avg_weight,
				lifespan, coalesce(details,''),
				coalesce(alternate_names,''),coalesce(geographic_origin,'')
				from dog_breeds where breed = ?`

	row := d.Db.QueryRowContext(ctx, query, b)
	var dogBreed DogBreed
	err := row.Scan(&dogBreed.Id,
		&dogBreed.Breed, &dogBreed.WeightLowLbs, &dogBreed.WeightHighLbs,
		&dogBreed.AverageWeight, &dogBreed.Lifespan, &dogBreed.Details,
		&dogBreed.AlternateNames, &dogBreed.GeographicOrigin)
	if err != nil {
		return nil, err
	}
	return &dogBreed, err
}

func (m *mysqlRepository) GetDogOfMonthById(id int) (*DogOfMonth, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `select id, video, image from dog_of_month where id = ?`

	row := m.Db.QueryRowContext(ctx, query, id)

	var dog DogOfMonth
	err := row.Scan(&dog.Id, &dog.Video, &dog.Image)
	if err != nil {
		return nil, err
	}
	return &dog, nil
}
