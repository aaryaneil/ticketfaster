package com.ticketfaster.orders.repository;


import com.ticketfaster.orders.model.Order;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface OrderRepository extends JpaRepository<Order, Long> {
    // You can add custom query methods here if needed
    // Example: List<Order> findByUserId(String userId);
}